package main

import (
	"pc-ziegert.de/media_service/common/config"
	"pc-ziegert.de/media_service/common/constant"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
	s "pc-ziegert.de/media_service/service"
	"pc-ziegert.de/media_service/service/utils"
)

func main() {
	conf := config.LoadConfig()
	img := m.NewImage("42", 1337)
	img.NodeType = "image/jpeg"

	fileName := "sample"
	fileExtension := "pdf"
	download_key := "0/" + fileName + "." + fileExtension
	upload_key := "media/" + fileName + ".jpg"

	s3Download, _ := utils.GenerateS3PresignDownloadUrl(nil, conf, download_key, "application/json")
	s3Upload, _, _ := utils.GenerateS3PresignUploadUrl(nil, conf, upload_key, "application/json")

	img.S3DownloadUrl = s3Download.URL
	c := m.NewConversion(100, 100, false)
	c.S3UploadUrl = s3Upload
	img.Conversions = append(img.Conversions, c)
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
	pub, err := mq.NewPublishing(img)
	mq.Publish(
		constant.RabbitMQExchangeName,
		constant.RabbitMQWorkerRoutingKey,
		false,
		false,
		*pub)
}
