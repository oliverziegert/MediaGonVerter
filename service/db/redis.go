package db

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"pc-ziegert.de/media_service/common/config"
	e "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	"pc-ziegert.de/media_service/service/db/repo"
)

type Redis struct {
	config *config.Config

	rdb *redis.Client
	//txm *tx.TransactionManager
	iRepo *repo.ImageRepo
	tRepo *repo.TenantRepo
}

// NewRedis creates a new Redis for the supplied configuration
func NewRedis(config *config.Config) *Redis {
	return &Redis{
		config: config,
		rdb:    nil,
	}
}

func (r *Redis) NewClient() {
	address := fmt.Sprintf("%s:%d", r.config.Redis.IP, r.config.Redis.Port)

	options := &redis.Options{
		Network: r.config.Redis.Network,
		Addr:    address,
		//Dialer:             nil,
		//OnConnect:          nil,
		Username: r.config.Redis.Username,
		Password: r.config.Redis.Password,
		DB:       int(r.config.Redis.DB),
		//MaxRetries:         0,
		//MinRetryBackoff:    0,
		//MaxRetryBackoff:    0,
		//DialTimeout:        0,
		//ReadTimeout:        0,
		//WriteTimeout:       0,
		//PoolFIFO:           false,
		//PoolSize:           0,
		//MinIdleConns:       0,
		//MaxConnAge:         0,
		//PoolTimeout:        0,
		//IdleTimeout:        0,
		//IdleCheckFrequency: 0,
		//TLSConfig:          nil,
		//Limiter:            nil,
	}
	r.rdb = redis.NewClient(options)
	_, err := r.rdb.Ping(r.rdb.Context()).Result()
	if err != nil {
		err := e.WrapError(e.SysRedisConnectionFailed, "Failed to connect to Redis", err)
		log.Fatal(err.StackTrace())
	}
}

// ClosePool closes the underlying database.
func (r *Redis) CloseClient() {
	if r.rdb == nil {
		return
	}

	err := r.rdb.Close()
	if err != nil {
		log.Fatalf("Could not close database connection!\nError: %s", err)
	}
}

// GetTransactionManager provides the TransactionManager.
//func (r *Redis) GetTransactionManager() *tx.TransactionManager {
//	if r.txm == nil {
//		r.txm = tx.NewTransactionManager(r.pool)
//	}
//
//	return r.txm
//}

// GetImageRepo provides the ImageRepo.
func (r *Redis) GetImageRepo() *repo.ImageRepo {
	if r.iRepo == nil {
		r.iRepo = repo.NewImageRepo(r.rdb, r.config.Data.ExpirationDuration)
	}
	return r.iRepo
}

func (r *Redis) GetTenantRepo() *repo.TenantRepo {
	if r.tRepo == nil {
		r.tRepo = repo.NewTenantRepo(r.rdb)
	}
	return r.tRepo
}
