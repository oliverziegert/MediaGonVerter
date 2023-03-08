package worker

import "pc-ziegert.de/media_service/common/constant"

func GetRabbitMQWorkerRoutingKeys() []string {
	return []string{
		constant.RabbitMQWorkerRoutingKeyPrefix + "image.webp",
		constant.RabbitMQWorkerRoutingKeyPrefix + "image.heic",
		constant.RabbitMQWorkerRoutingKeyPrefix + "image.jpeg",
		constant.RabbitMQWorkerRoutingKeyPrefix + "image.svg+xml",
		constant.RabbitMQWorkerRoutingKeyPrefix + "image.jp2",
		constant.RabbitMQWorkerRoutingKeyPrefix + "image.gif",
		constant.RabbitMQWorkerRoutingKeyPrefix + "image.png",
		constant.RabbitMQWorkerRoutingKeyPrefix + "image.tiff",
		constant.RabbitMQWorkerRoutingKeyPrefix + "image.heif",
		constant.RabbitMQWorkerRoutingKeyPrefix + "application.pdf",
	}
}
