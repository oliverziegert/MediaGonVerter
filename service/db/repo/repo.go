package repo

import (
	"github.com/go-redis/redis/v8"
)

type repo struct {
	rdb *redis.Client
}
