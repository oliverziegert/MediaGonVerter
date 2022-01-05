package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"pc-ziegert.de/media_service/config"
	"pc-ziegert.de/media_service/constant"
	"pc-ziegert.de/media_service/log"
	"syscall"
	"time"
)

func main() {
	// Load config
	conf := config.LoadConfig()

	// Set logging level
	log.SetLevel(conf.Log.Level)

	log.Infof("Start Media Service %s.", constant.AppVersion)

	// Create initializer
	init := NewInitializer(conf)

	// Open and create/update database
	// db := init.GetDb()
	// db.OpenDb()
	// defer db.CloseDb()
	// db.UpdateDb()

	// Schedule jobs
	init.GetJobService().ScheduleJobs()

	// Create router
	router := gin.Default()

	// Set release mode for router
	if init.conf.Log.Level != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	// configure router
	configureRouter(init, router)

	// Add Root handlers
	addRootHandlers(router)

	// Add Swagger UI handlers
	addSwaggerUiHandlers(router)

	// Configure API routing
	configureApiRouting(init, router)

	// Create server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", init.conf.HTTP.Port),
		Handler:        router,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}

	// Start HTTP server
	go func() {
		// service connections
		log.Infof("start server: :%d", init.conf.HTTP.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Info("timeout of 5 seconds.")
	}
	log.Info("Server exiting")
}

func configureRouter(init *Initializer, router *gin.Engine) {
	if init.conf.HTTP.Gzip.Enabled {
		router.Use(gzip.Gzip(gzip.BestCompression))
	}
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = init.conf.HTTP.Origin.All.Enabled
	//corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(corsConfig))
}

func configureApiRouting(init *Initializer, router *gin.Engine) {
	// Create API sub route
	ar := router.Group(constant.ApiPath)

	// Create protected middleware route
	ar.Use(init.GetTransactionApiMiddleware().ServeHTTP())
	ar.Use(init.GetErrorApiMiddleware().ServeHTTP())
	ar.Use(init.GetSecurityApiMiddleware().ServeHTTP())
	ar.Use(init.GetAuthCheckApiMiddleware().ServeHTTP())

	// Get controllers
	imgCtrl := init.GetImageApiController()

	// Add protected endpoints
	ar.GET("/image/:token/:size", imgCtrl.GetImageHandler())
	ar.HEAD("/image/:token/:size", imgCtrl.HeadImageHandler())
}

func addSwaggerUiHandlers(router *gin.Engine) {
	router.GET(constant.ApiPath, func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%s/swagger-ui/index.html", constant.ApiPath))
	})
	router.StaticFS(fmt.Sprintf("%s/swagger-ui/", constant.ApiPath), http.Dir("./swagger-ui"))
}

func addRootHandlers(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, constant.ApiPath)
	})
}
