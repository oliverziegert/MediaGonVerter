package repo

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	error2 "pc-ziegert.de/media_service/common/error"
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
func (i *TenantRepo) ExistsTenant(ctx *gin.Context, t *model.Tenant) (bool, *error2.Error) {
	result, err := i.rdb.Exists(ctx, t.GetKeyString()).Result()
	if err != nil {
		return false, error2.WrapError(error2.ValIdInvalid, "ToDo: Error handling", err)
	}
	if result > 0 {
		return true, nil
	}
	return false, nil
}

// GetTenant retrieves a tenant by primary key.
func (i *TenantRepo) GetTenant(ctx *gin.Context, t *model.Tenant) *error2.Error {
	valueMap, err := i.rdb.HGetAll(ctx, t.GetKeyString()).Result()
	if err != nil {
		return error2.WrapError(1, "ToDo: Error handling", err)
	}
	err = t.MapToTenant(valueMap)
	if err != nil {
		return error2.WrapError(1, "ToDo: Error handling", err)
	}
	return nil
}

// SetImage creates a new image
func (i *TenantRepo) SetTenant(ctx *gin.Context, t *model.Tenant, ex time.Duration) *error2.Error {
	err := i.rdb.HSet(ctx, t.GetKeyString(), t.GetValueMap()).Err()
	if err != nil {
		return error2.WrapError(1, "ToDo: Error handling", err)
	}
	err = i.rdb.Expire(ctx, t.GetKeyString(), ex).Err()
	if err != nil {
		return error2.WrapError(1, "ToDo: Error handling", err)
	}
	return nil
}
