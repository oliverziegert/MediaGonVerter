package repo

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	"pc-ziegert.de/media_service/common/model"
	"time"
)

type TenantRepo struct {
	repo
}

func NewTenantRepo(rdb *redis.Client) *TenantRepo {
	return &TenantRepo{
		repo: repo{rdb: rdb},
	}
}

// ExistsTenant checks if tenant exists
func (i *TenantRepo) ExistsTenant(ctx *gin.Context, t *model.Tenant) (bool, *e.Error) {
	result, err := i.rdb.Exists(ctx, t.GetKeyString()).Result()
	if err != nil {
		return false, e.WrapError(e.ValIdInvalid, "ToDo: Error handling", err)
	}
	if result > 0 {
		return true, nil
	}
	return false, nil
}

// GetTenant retrieves a tenant by primary key.
func (i *TenantRepo) GetTenant(ctx *gin.Context, t *model.Tenant) *e.Error {
	valueMap, err := i.rdb.HGetAll(ctx, t.GetKeyString()).Result()
	if err != nil {
		err := e.WrapError(1, "ToDo: Error handling", err)
		log.Debug(err.StackTrace())
		return err
	}
	newErr := t.MapToTenant(valueMap)
	if newErr != nil {
		err := e.WrapError(1, "ToDo: Error handling", newErr)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}

// GetTenantTtl retrieves tts of a specific tenant by primary key if present.
func (i *TenantRepo) GetTenantTtl(ctx *gin.Context, t *model.Tenant) (*time.Duration, *e.Error) {
	// check if tenant is cached
	exists, err := i.ExistsTenant(ctx, t)
	if err != nil {
		err := e.WrapError(1, "ToDo: Error handling", err)
		log.Debug(err.StackTrace())
		return nil, err
	}

	// if not return duration of 0
	if !exists {
		var t time.Duration = 0
		return &t, nil
	}

	// get tll
	ttl, errExt := i.rdb.TTL(ctx, t.GetKeyString()).Result()
	if errExt != nil {
		err := e.WrapError(1, "ToDo: Error handling", errExt)
		log.Debug(err.StackTrace())
		return nil, err
	}

	return &ttl, nil
}

// SetTenant creates or updates a tenant
func (i *TenantRepo) SetTenant(ctx *gin.Context, t *model.Tenant, ex *time.Duration) *e.Error {
	err := i.rdb.HSet(ctx, t.GetKeyString(), t.GetValueMap()).Err()
	if err != nil {
		return e.WrapError(1, "ToDo: Error handling", err)
	}
	err = i.rdb.Expire(ctx, t.GetKeyString(), *ex).Err()
	if err != nil {
		return e.WrapError(1, "ToDo: Error handling", err)
	}
	return nil
}

// setTenantTtl sets ttl of a tenant
func (i *TenantRepo) setTenantTtl(ctx *gin.Context, t *model.Tenant, ex *time.Duration) *e.Error {
	// check if tenant is cached
	exists, err := i.ExistsTenant(ctx, t)
	if err != nil {
		err := e.WrapError(1, "ToDo: Error handling", err)
		log.Debug(err.StackTrace())
		return err
	}

	// if not return
	if !exists {
		err := e.NewError(1, "ToDo: Error handling")
		log.Debug(err.StackTrace())
		return err
	}

	errExt := i.rdb.Expire(ctx, t.GetKeyString(), *ex).Err()
	if errExt != nil {
		err := e.WrapError(1, "ToDo: Error handling", errExt)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}

// BackoffTenant increases ttl exponentially of a tenant
func (i *TenantRepo) BackoffTenant(ctx *gin.Context, t *model.Tenant) *e.Error {
	ttl, err := i.GetTenantTtl(ctx, t)
	if err != nil {
		err := e.WrapError(1, "ToDo: Error handling", err)
		log.Debug(err.StackTrace())
		return err
	}

	if ttl.Microseconds() == 0 {
		err := e.NewError(1, "Tenant does not exist, cant increase ttl")
		log.Debug(err.StackTrace())
		return err
	}

	ttlNew := *ttl + constant.TenantBackoffDuration
	if ttlNew > constant.TenantMaxDuration {
		ttlNew = constant.TenantMaxDuration
	}

	err = i.setTenantTtl(ctx, t, &ttlNew)
	if err != nil {
		err := e.WrapError(1, "ToDo: Error handling", err)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}
