package worker

func GetRabbitMQWorkerRoutingKeys() []string {
	return []string{
		"worker.image.webp",
		"worker.image.heic",
		"worker.image.jpeg",
		"worker.image.svg+xml",
		"worker.image.jp2",
		"worker.image.gif",
		"worker.image.png",
		"worker.image.tiff",
		"worker.image.heif",
		"worker.application.pdf",
	}
}
