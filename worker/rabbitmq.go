package worker

import (
	"github.com/streadway/amqp"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
	"pc-ziegert.de/media_service/service/mq"
	"pc-ziegert.de/media_service/service/utils"
	"sync"
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

func Publish(wg *sync.WaitGroup, r *mq.RabbitMQ, cRes <-chan m.Image) {
	defer wg.Done()

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
