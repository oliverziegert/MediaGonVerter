package mq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"pc-ziegert.de/media_service/common/config"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	"pc-ziegert.de/media_service/common/model"
)

type RabbitMQ struct {
	config *config.Config
	con    *Connection
	ch     *Channel
}

func NewRabbtmq(config *config.Config) *RabbitMQ {
	return &RabbitMQ{
		config: config,
		con:    nil,
		ch:     nil,
	}
}

func (r *RabbitMQ) OpenRabbitmq() *e.Error {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		r.config.RabbitMQ.Username,
		r.config.RabbitMQ.Password,
		r.config.RabbitMQ.IP,
		r.config.RabbitMQ.Port,
		r.config.RabbitMQ.VirtualHost)
	con, err := Dial(url)
	if err != nil {
		err := e.WrapError(e.SysRabbitConnectionFailed, "Failed to connect to RabbitMQ.", err)
		log.Debug(err.StackTrace())
		return err
	}
	r.con = con

	ch, err := r.con.Channel()
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to open a RabbitMQ channel", err)
		log.Debug(err.StackTrace())
		return err
	}
	r.ch = ch

	return nil
}

func (r *RabbitMQ) CloseRabbitmq() *e.Error {
	err := r.ch.Close()
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to close a RabbitMQ channel", err)
		log.Debug(err.StackTrace())
		return err
	}

	err = r.con.Close()
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to close a RabbitMQ connection", err)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}

func (r *RabbitMQ) DeclareExchange(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args amqp.Table) *e.Error {
	err := r.ch.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to declare an exchange", err)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}

func (r *RabbitMQ) DeclareQueue(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (amqp.Queue, *e.Error) {
	q, err := r.ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to declare a queue", err)
		log.Debug(err.StackTrace())
		return q, err
	}
	return q, nil
}

func (r *RabbitMQ) QueueBind(name string, key string, exchange string, noWait bool, args amqp.Table) *e.Error {
	err := r.ch.QueueBind(name, key, exchange, noWait, args)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to declare a queue", err)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}

func (r *RabbitMQ) Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table) (<-chan amqp.Delivery, *e.Error) {
	msgs, err := r.ch.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return msgs, err
	}
	return msgs, nil
}

func (r *RabbitMQ) NewPublishing(image *model.Image) (*amqp.Publishing, *e.Error) {
	json, err := image.JSON()
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return nil, err
	}
	return &amqp.Publishing{
		Headers:         nil,
		ContentType:     "application/json",
		ContentEncoding: "utf-8",
		DeliveryMode:    2,
		CorrelationId:   image.GetKeyString(),
		Body:            json,
		Type:            constant.RabbitMQImageMessageType,
	}, nil
}

func (r *RabbitMQ) Publish(exchange, key string, mandatory bool, immediate bool, msg amqp.Publishing) *e.Error {
	err := r.ch.PublishWithContext(context.TODO(), exchange, key, mandatory, immediate, msg)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}

func (r *RabbitMQ) Configure() *e.Error {
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
		log.Debug(err.StackTrace())
		return err
	}
	_, err = r.DeclareQueue(
		constant.RabbitMQQueueMediaServiceName,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return err
	}
	err = r.QueueBind(
		constant.RabbitMQQueueMediaServiceName,
		constant.RabbitMQMediaServiceRoutingKey,
		constant.RabbitMQExchangeName,
		false,
		nil)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}
