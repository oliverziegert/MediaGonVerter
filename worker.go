package main

import (
	"pc-ziegert.de/media_service/common/config"
	"pc-ziegert.de/media_service/common/constant"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
	s "pc-ziegert.de/media_service/service"
	w "pc-ziegert.de/media_service/worker"
	"sync"
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

	msg, err := w.Consume(mq)
	if err != nil {
		l.Debug(err.StackTrace())
		l.Error(err.Error())
		panic(err)
	}

	var wg sync.WaitGroup
	resp := make(chan m.Image)

	wg.Add(1)
	go w.Publish(&wg, mq, resp)

	w.Convert(<-msg, resp)

	wg.Wait()

	l.Info("Worker exiting")
}
