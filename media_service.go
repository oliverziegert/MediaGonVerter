package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"net/http"
	"os"
	"os/signal"
	"pc-ziegert.de/media_service/common/config"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	s "pc-ziegert.de/media_service/service"
	"pc-ziegert.de/media_service/service/mq"
	"pc-ziegert.de/media_service/service/utils"
	"syscall"
	"time"
)

func main() {
	// Load config
	conf := config.LoadConfig()

	// Set logging level
	l.SetLevel(conf.Log.Level)

	l.Infof("Start Media Service %s.", constant.AppVersion)

	// Create initializer
	i := s.NewInitializer(conf)

	// Open RabbitMQ connection
	// create Exchange and routing keys
	mq := i.GetRabbitMQ()
	mq.OpenRabbitmq()
	defer mq.CloseRabbitmq()
	mq.Configure()

	// Open and create/update database
	redis := i.GetRedis()
	redis.NewClient()
	defer redis.CloseClient()

	// Schedule jobs
	i.GetJobService().ScheduleJobs()

	// Create router
	router := gin.Default()

	// Set release mode for router
	if i.GetConfig().Log.Level != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	// configure router
	s.ConfigureRouter(i, router)

	// Add Root handlers
	s.AddRootHandlers(router)

	// Add probe handlers
	s.AddProbeHandlers(i, router)

	// Add probe handlers
	s.AddMetricsHandlers(i, router)

	// Add Swagger UI handlers
	s.AddSwaggerUiHandlers(router)

	// Configure API routing
	s.ConfigureApiRouting(i, router)

	// Create server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", i.GetConfig().HTTP.Port),
		Handler:        router,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}

	// Start HTTP server
	go func() {
		// s connections
		l.Infof("start server: :%d", i.GetConfig().HTTP.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatalf("listen: %s", err)
		}
	}()
	msgs, err := consume(mq)
	if err != nil {
		l.Error("Can not consume queue")
	}

	rabServ := i.GetRabbitMqService()
	go rabServ.Consumer(msgs)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	l.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		l.Fatalf("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		l.Info("timeout of 5 seconds.")
	}
	l.Info("Server exiting")
}

func consume(r *mq.RabbitMQ) (<-chan amqp.Delivery, *e.Error) {
	h := utils.GetHostname()
	return r.Consume(
		constant.RabbitMQQueueMediaServiceName,
		*h,
		false,
		false,
		false,
		false,
		nil,
	)
}
