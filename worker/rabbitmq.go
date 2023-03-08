package worker

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
	"pc-ziegert.de/media_service/service/mq"
	"pc-ziegert.de/media_service/service/utils"
)

func Consume(r *mq.RabbitMQ) (<-chan amqp.Delivery, *e.Error) {
	h := utils.GetHostname()
	return r.Consume(
		constant.RabbitMQQueueWorkerName,
		*h,
		false,
		false,
		false,
		false,
		nil,
	)
}

func Publish(r *mq.RabbitMQ, cRes <-chan m.Image) {
	for i := range cRes {
		err := publish(r, &i)
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, "", err)
			l.Debug(err.StackTrace())
		}
	}
}

func publish(r *mq.RabbitMQ, i *m.Image) *e.Error {
	pub, err := r.NewPublishing(i)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "", err)
		l.Debug(err.StackTrace())
		return err
	}
	err = r.Publish(
		constant.RabbitMQExchangeName,
		constant.RabbitMQMediaServiceRoutingKey,
		false,
		false,
		*pub,
	)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "", err)
		l.Debug(err.StackTrace())
		return err
	}
	return nil
}

func ConfigureRabbitMq(r *mq.RabbitMQ) *e.Error {
	err := r.DeclareExchange(
		constant.RabbitMQExchangeName,
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		l.Debug(err.StackTrace())
		return err
	}
	_, err = r.DeclareQueue(
		constant.RabbitMQQueueWorkerName,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		l.Debug(err.StackTrace())
		return err
	}
	for _, routingKey := range GetRabbitMQWorkerRoutingKeys() {
		err = r.QueueBind(
			constant.RabbitMQQueueWorkerName,
			routingKey,
			constant.RabbitMQExchangeName,
			false,
			nil)
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
			l.Debug(err.StackTrace())
			return err
		}
	}
	return nil
}
