package service

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	c "pc-ziegert.de/media_service/common/config"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
	"pc-ziegert.de/media_service/service/db/repo"
	"time"
)

type RabbitMqService struct {
	conf  *c.Config
	iRepo *repo.ImageRepo
}

func NewRabbitMQService(conf *c.Config, iRepo *repo.ImageRepo) *RabbitMqService {
	return &RabbitMqService{
		conf:  conf,
		iRepo: iRepo,
	}
}

func (rServ *RabbitMqService) Consumer(msgs <-chan amqp.Delivery) {

	for msg := range msgs {
		rServ.consume(msg, &rServ.conf.Data.ExpirationDuration)
	}
}

func (rServ *RabbitMqService) consume(msg amqp.Delivery, ex *time.Duration) *e.Error {
	defer msg.Ack(false)
	switch msg.Type {
	case constant.RabbitMQImageMessageType:
		{
			i := m.Image{}
			err := i.JsonToImage(msg.Body)
			if err != nil {
				err := e.WrapError(e.ValIdInvalid, "", err)
				l.Debug(err.StackTrace())
				return err
			}

			exists, err := rServ.iRepo.ExistsImage(context.Background(), &i)
			if err != nil {
				err := e.WrapError(e.ValIdInvalid, "", err)
				l.Debug(err.StackTrace())
				return err
			}

			if !exists {
				err := rServ.iRepo.SetImage(context.Background(), &i, ex)
				if err != nil {
					err := e.WrapError(e.ValIdInvalid, "", err)
					l.Debug(err.StackTrace())
					return err

				}
				return nil
			}

			err = rServ.iRepo.SetImageTtl(context.Background(), &i, ex)
			if err != nil {
				err := e.WrapError(e.ValIdInvalid, "", err)
				l.Debug(err.StackTrace())
				return err
			}

			err = rServ.iRepo.UpdateImage(context.Background(), &i)
			if err != nil {
				err := e.WrapError(e.ValIdInvalid, "", err)
				l.Debug(err.StackTrace())
				return err
			}
			return nil
		}
	default:
		err := e.NewError(e.ValIdInvalid, fmt.Sprintf("unknown message type: %s", msg.Type))
		l.Debug(err.StackTrace())
		msg.Ack(false)
		return err
	}
}
