package service

import (
	"github.com/gin-gonic/gin"
	"pc-ziegert.de/media_service/common/config"
	"pc-ziegert.de/media_service/common/constant"
	error2 "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	"pc-ziegert.de/media_service/common/model"
	"pc-ziegert.de/media_service/service/db/repo"
	"pc-ziegert.de/media_service/service/mq"
)

// TenantService contains image related logic.
type TenantService struct {
	conf  *config.Config
	tRepo *repo.TenantRepo
	mq    *mq.RabbitMQ
}

// NewTenantService NewUserService create a new user service.
func NewTenantService(conf *config.Config, tRepo *repo.TenantRepo, mq *mq.RabbitMQ) *TenantService {
	return &TenantService{
		conf:  conf,
		tRepo: tRepo,
		mq:    mq,
	}
}

func (tServ *TenantService) TenantState(ctx *gin.Context) (*model.TenantState, *error2.Error) {
	tState := model.TenantStateUnknown
	value, exists := ctx.Get(constant.ContextKeyJWTTokenClaims)
	if !exists {
		err := error2.NewError(error2.ValIdInvalid, "")
		log.Debug(err.StackTrace())
		return nil, err
	}

	tUuid := value.(*CustomClaims).TenantUuid
	t := model.NewTenant(tUuid)

	exists, err := tServ.tRepo.ExistsTenant(ctx, t)
	if err != nil {
		err := error2.WrapError(error2.ValIdInvalid, "", err)
		log.Debug(err.StackTrace())
		return nil, err
	}
	if !exists {
		return &tState, nil
	}

	err = tServ.tRepo.GetTenant(ctx, t)
	if err != nil {
		err := error2.WrapError(error2.ValIdInvalid, "", err)
		log.Debug(err.StackTrace())
		return nil, err
	}

	return &t.TenantState, nil
}
