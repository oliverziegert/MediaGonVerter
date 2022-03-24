package repo

import (
	"context"
	"github.com/go-redis/redis/v8"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
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
func (i *ImageRepo) GetImage(ctx context.Context, image *m.Image) *e.Error {
	valueMap, err := i.rdb.HGetAll(ctx, image.GetKeyString()).Result()
	if err != nil {
		return e.WrapError(1, "ToDo: Error handling", err)
	}
	newErr := image.MapToImage(valueMap)
	if newErr != nil {
		return e.WrapError(1, "ToDo: Error handling", err)
	}
	return nil
}

// SetImage creates a new image
func (i *ImageRepo) SetImage(ctx context.Context, image *m.Image, ex *time.Duration) *e.Error {
	err := i.rdb.HSet(ctx, image.GetKeyString(), image.GetValueMap()).Err()
	if err != nil {
		return e.WrapError(1, "ToDo: Error handling", err)
	}
	err = i.rdb.Expire(ctx, image.GetKeyString(), *ex).Err()
	if err != nil {
		return e.WrapError(1, "ToDo: Error handling", err)
	}
	return nil
}

// UpdateImage Updates a existing image
func (i *ImageRepo) UpdateImage(ctx context.Context, image *m.Image) *e.Error {
	cachedImage := *image
	err := i.GetImage(ctx, &cachedImage)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "", err)
		l.Debug(err.StackTrace())
		return err
	}

	if image.ServiceName != cachedImage.ServiceName {
		vm := make(map[string]string)
		image.GetValueMapForValue(&vm, "serviceName")
		err := i.rdb.HSet(ctx, image.GetKeyString(), vm).Err()
		if err != nil {
			newErr := e.WrapError(e.ValIdInvalid, "", err)
			l.Debug(newErr.StackTrace())
			return newErr
		}
	}

	if image.NodeType != cachedImage.NodeType {
		vm := make(map[string]string)
		image.GetValueMapForValue(&vm, "nodeType")
		err := i.rdb.HSet(ctx, image.GetKeyString(), vm).Err()
		if err != nil {
			newErr := e.WrapError(e.ValIdInvalid, "", err)
			l.Debug(newErr.StackTrace())
			return newErr
		}
	}

	if image.NodeSize != cachedImage.NodeSize {
		vm := make(map[string]string)
		image.GetValueMapForValue(&vm, "nodeSize")
		err := i.rdb.HSet(ctx, image.GetKeyString(), vm).Err()
		if err != nil {
			newErr := e.WrapError(e.ValIdInvalid, "", err)
			l.Debug(newErr.StackTrace())
			return newErr
		}
	}

	for _, c := range image.Conversions {
		cc, err := cachedImage.GetConversionBySize(c.Width, c.Height, c.Crop)
		if err != nil || c.State > cc.State || c.Expires.Unix() > cc.Expires.Unix() {
			vm := make(map[string]string)
			c.GetValueMap(&vm)
			err := i.rdb.HSet(ctx, image.GetKeyString(), vm).Err()
			if err != nil {
				newErr := e.WrapError(e.ValIdInvalid, "", err)
				l.Debug(newErr.StackTrace())
				return newErr
			}
		}
	}

	return nil
}

// ExistsImage checks if image exists
func (i *ImageRepo) ExistsImage(ctx context.Context, image *m.Image) (bool, *e.Error) {
	result, err := i.rdb.Exists(ctx, image.GetKeyString()).Result()
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "", err)
		l.Debug(err.StackTrace())
		return false, nil
	}
	if result > 0 {
		return true, nil
	}
	return false, nil
}

// DeleteImage deletes a new image
func (i *ImageRepo) DeleteImage(ctx context.Context, image *m.Image) *e.Error {
	err := i.rdb.Del(ctx, image.GetKeyString()).Err()
	if err != nil {
		err := e.WrapError(1, "ToDo: Error handling", err)
		l.Debug(err.StackTrace())
		return err
	}
	return nil
}

// SetImageTtl sets ttl of a image
func (i *ImageRepo) SetImageTtl(ctx context.Context, image *m.Image, ex *time.Duration) *e.Error {
	// check if tenant is cached
	exists, err := i.ExistsImage(ctx, image)
	if err != nil {
		err := e.WrapError(1, "ToDo: Error handling", err)
		l.Debug(err.StackTrace())
		return err
	}

	// if not return
	if !exists {
		err := e.NewError(1, "ToDo: Error handling")
		l.Debug(err.StackTrace())
		return err
	}

	errExt := i.rdb.Expire(ctx, image.GetKeyString(), *ex).Err()
	if errExt != nil {
		err := e.WrapError(1, "ToDo: Error handling", errExt)
		l.Debug(err.StackTrace())
		return err
	}
	return nil
}
