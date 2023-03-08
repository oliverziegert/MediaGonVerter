package main

import (
	"github.com/davidbyttow/govips/v2/vips"
	"pc-ziegert.de/media_service/common/config"
	"pc-ziegert.de/media_service/common/constant"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
	s "pc-ziegert.de/media_service/service"
	w "pc-ziegert.de/media_service/worker"
)

func main() {
	// Load config
	conf := config.LoadConfig()

	// Set logging level
	l.SetLevel(conf.Log.Level)

	l.Infof("Start Media Service Worker %s.", constant.AppVersion)

	// Create initializer
	init := s.NewInitializer(conf)

	// Open RabbitMQ connection
	// create Exchange and routing keys
	mq := init.GetRabbitMQ()
	err := mq.OpenRabbitmq()
	if err != nil {
		l.Debug(err.StackTrace())
		l.Error(err.Error())
	}
	defer mq.CloseRabbitmq()
	err = mq.Configure()
	if err != nil {
		l.Debug(err.StackTrace())
		l.Error(err.Error())
	}

	msgs, err := w.Consume(mq)
	if err != nil {
		l.Debug(err.StackTrace())
		l.Error(err.Error())
		panic(err)
	}

	resp := make(chan m.Image)
	var forever chan struct{}

	go w.Publish(mq, resp)

	vips.LoggingSettings(w.VipsLogger, vips.LogLevelInfo)
	vips.Startup(nil)
	defer vips.Shutdown()

	go func() {
		for msg := range msgs {
			w.Convert(msg, resp)
		}
	}()

	l.Info("Waiting for logs. To exit press CTRL+C")
	<-forever

}
