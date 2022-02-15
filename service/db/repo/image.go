package repo

import (
	"context"
	"github.com/go-redis/redis/v8"
	e "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/model"
	"time"
)

type ImageRepo struct {
	repo
	ex time.Duration
}

func NewImageRepo(rdb *redis.Client, ex time.Duration) *ImageRepo {
	return &ImageRepo{
		repo: repo{rdb: rdb},
		ex:   ex,
	}
}

// --- Image functions ---

// GetImages retrieves all images.
func (i *ImageRepo) GetImages(ctx context.Context) *e.Error {
	return nil
}

// GetImage retrieves an image by primary key.
func (i *ImageRepo) GetImage(ctx context.Context, image *model.Image) *e.Error {
	valueMap, err := i.rdb.HGetAll(ctx, image.GetKeyString()).Result()
	if err != nil {
		return e.WrapError(1, "ToDo: Error handling", err)
	}
	err = image.MapToImage(valueMap)
	if err != nil {
		return e.WrapError(1, "ToDo: Error handling", err)
	}
	return nil
}

// SetImage creates a new image
func (i *ImageRepo) SetImage(ctx context.Context, image *model.Image) *e.Error {
	err := i.rdb.HSet(ctx, image.GetKeyString(), image.GetValueMap()).Err()
	if err != nil {
		return e.WrapError(1, "ToDo: Error handling", err)
	}
	err = i.rdb.Expire(ctx, image.GetKeyString(), i.ex).Err()
	if err != nil {
		return e.WrapError(1, "ToDo: Error handling", err)
	}
	return nil
}

// UpdateImage Updates a existing image
func (i *ImageRepo) UpdateImage(ctx context.Context, image *model.Image) *e.Error {
	// ToDo: only diff update
	i.rdb.HSet(ctx, image.GetKeyString(), image.GetValueMap())
	return nil
}

// ExistsImage checks if image exists
func (i *ImageRepo) ExistsImage(ctx context.Context, image *model.Image) (bool, *e.Error) {
	if result, err := i.rdb.Exists(ctx, image.GetKeyString()).Result(); err == nil && result > 0 {
		return true, nil
	} else if err == nil && result <= 0 {
		return false, nil
	} else {
		return false, e.WrapError(1, "ToDo: Error handling", err)
	}
}

// DeleteImage deletes a new image
func (i *ImageRepo) DeleteImage(ctx context.Context, image *model.Image) *e.Error {
	return nil
}
