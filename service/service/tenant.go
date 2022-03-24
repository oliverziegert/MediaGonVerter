package service

import (
	"github.com/gin-gonic/gin"
	"pc-ziegert.de/media_service/common/config"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
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

func (tServ *TenantService) TenantState(ctx *gin.Context) (*m.TenantState, *e.Error) {
	tState := m.TenantStateUnknown

	// Get jwt token claims
	cc, exists := GetJwtClaimsFromContext(ctx)
	if !exists {
		err := e.NewError(e.ValIdInvalid, "Failed to register a consumer.")
		l.Debug(err.StackTrace())
		return nil, err
	}

	// create tenant object from jwt token claim uuid
	t := m.NewTenant(cc.TenantUuid)

	// check if tenant is cached
	exists, err := tServ.tRepo.ExistsTenant(ctx, t)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "", err)
		l.Debug(err.StackTrace())
		return nil, err
	}

	// if tenant is not known
	if !exists {
		return &tState, nil
	}

	err = tServ.tRepo.GetTenant(ctx, t)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "", err)
		l.Debug(err.StackTrace())
		return nil, err
	}

	return &t.TenantState, nil
}

func (tServ *TenantService) BackoffTenant(ctx *gin.Context) *e.Error {
	// Get jwt token claims
	cc, exists := GetJwtClaimsFromContext(ctx)
	if !exists {
		err := e.NewError(e.ValIdInvalid, "Failed to register a consumer.")
		l.Debug(err.StackTrace())
		return err
	}

	// create tenant object from jwt token claim uuid
	t := m.NewTenant(cc.TenantUuid)

	// check if tenant is cached
	exists, err := tServ.tRepo.ExistsTenant(ctx, t)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "", err)
		l.Debug(err.StackTrace())
		return err
	}

	t.TenantState = m.TenantStateInvalid
	if !exists {
		ibd := constant.TenantInitialBackoffDuration
		err = tServ.tRepo.SetTenant(ctx, t, &ibd)
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, "", err)
			l.Debug(err.StackTrace())
			return err
		}
	}

	err = tServ.tRepo.BackoffTenant(ctx, t)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "", err)
		l.Debug(err.StackTrace())
		return err
	}

	return nil
}
