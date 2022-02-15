package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"pc-ziegert.de/media_service/common/config"
	"pc-ziegert.de/media_service/common/constant"
	"pc-ziegert.de/media_service/common/log"
	"pc-ziegert.de/media_service/common/model"
	"pc-ziegert.de/media_service/service"
	"sync"
)

func main() {
	// Load config
	conf := config.LoadConfig()

	// Set logging level
	log.SetLevel(conf.Log.Level)

	log.Infof("Start Media Service %s.", constant.AppVersion)

	// Create initializer
	init := service.NewInitializer(conf)

	// Open RabbitMQ connection
	// create Exchange and routing keys
	mq := init.GetRabbitMQ()
	mq.OpenRabbitmq()
	defer mq.CloseRabbitmq()
	mq.Configure()

	// Open and create/update database
	redis := init.GetRedis()
	redis.NewClient()
	defer redis.CloseClient()

	// Schedule jobs
	init.GetJobService().ScheduleJobs()

	msgs, err := mq.Consume(
		constant.RabbitMQQueueWorkerName,
		"worker",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	convert(<-msgs)

	log.Info("Server exiting")
}

func convert(in amqp.Delivery) {
	var wg sync.WaitGroup

	defer in.Ack(false)

	i := model.Image{}
	i.JsonToImage(in.Body)

	for _, c := range i.Conversions {
		wg.Add(1)
		go convertFile(&wg, &i, c)
	}

	wg.Wait()
}

func convertFile(wg *sync.WaitGroup, i *model.Image, c *model.Conversion) {
	// call done when finished
	defer wg.Done()

	tds := fmt.Sprintf("./%v/%v-%d", "worker", i.TenantId, i.NodeId)

	err := os.Mkdir(tds, 750)
	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := http.Get(i.S3DownloadUrl.String())
	if err != nil {
		log.Fatal(err.Error())
	}
	if resp.StatusCode > 299 {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	newFilePath := fmt.Sprintf("%v/%dx%d", tds, c.Width, c.Height)
	outFile := openOrCreate(newFilePath)
	defer outFile.Close()

	// decode
	imageData, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatal("error decoding image")
	}

	// encode in new type
	switch i.NodeType {
	case "image/jpeg":
		err := jpeg.Encode(outFile, imageData, nil)
		if err != nil {
			log.Fatal("error converting to jpeg")
		}
	case "image/png":
		err := png.Encode(outFile, imageData)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "image/gif":
		err := gif.Encode(outFile, imageData, nil)
		if err != nil {
			log.Fatal("error converting to gif")
		}
	case "image/bmp":
		err := bmp.Encode(outFile, imageData)
		if err != nil {
			log.Fatal("error converting to bmp")
		}
	case "image/tiff":
		err := tiff.Encode(outFile, imageData, nil)
		if err != nil {
			log.Fatal("error converting to tiff")
		}
	}
}

func openOrCreate(filename string) *os.File {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal("error creating output file")
		}
		return file
	} else {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal("error opening output file")
		}
		return file
	}
}
