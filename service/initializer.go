package service

import (
	"math/rand"
	"pc-ziegert.de/media_service/common/config"
	ac "pc-ziegert.de/media_service/service/api/controller"
	am "pc-ziegert.de/media_service/service/api/middleware"
	"pc-ziegert.de/media_service/service/db"
	"pc-ziegert.de/media_service/service/mq"
	"pc-ziegert.de/media_service/service/service"
	"time"
)

// Initializer initializes application components.
type Initializer struct {
	conf *config.Config

	mq    *mq.RabbitMQ
	redis *db.Redis

	jobServ *service.JobService
	imgServ *service.ImageService
	jwtServ *service.JWTService
	tenServ *service.TenantService
	rabServ *service.RabbitMqService

	imgACtrl *ac.ImageController
	prbACtrl *ac.ProbeController
	metACtrl *ac.MetricsController

	traAMdw   *am.TransactionMiddleware
	errAMdw   *am.ErrorMiddleware
	secAMdw   *am.SecurityMiddleware
	authAMdw  *am.AuthCheckMiddleware
	tokenAMdw *am.TokenCheckMiddleware
}

// NewInitializer creates a new initializer.
func NewInitializer(conf *config.Config) *Initializer {
	rand.Seed(time.Now().UnixNano())
	return &Initializer{conf: conf}
}

// --- Config functions ---

// GetConfig returns an initialized config object.
func (i *Initializer) GetConfig() *config.Config {
	return i.conf
}

// --- RabbitMQ functions ---

// GetRabbitMQ returns an initialized RabbitMQ object.
func (i *Initializer) GetRabbitMQ() *mq.RabbitMQ {
	if i.mq == nil {
		i.mq = mq.NewRabbtmq(i.conf)
	}
	return i.mq
}

// --- Database functions ---

// GetRedis returns an initialized Redis object.
func (i *Initializer) GetRedis() *db.Redis {
	if i.redis == nil {
		i.redis = db.NewRedis(i.conf)
	}
	return i.redis
}

// --- Service functions ---

// GetJobService returns an initialized job service object
func (i *Initializer) GetJobService() *service.JobService {
	if i.jobServ == nil {
		i.jobServ = service.NewJobService()
	}
	return i.jobServ
}

// GetImageService returns an initialized image service object.
func (i *Initializer) GetImageService() *service.ImageService {
	if i.imgServ == nil {
		i.imgServ = service.NewImageService(
			i.conf,
			i.GetRedis().GetImageRepo(),
			i.mq)
	}
	return i.imgServ
}

// GetJWTService returns an initialized jwt service object.
func (i *Initializer) GetJWTService() *service.JWTService {
	if i.jwtServ == nil {
		i.jwtServ = service.NewJWTService(
			i.conf.Auth.Jwt.PublicKey)
	}
	return i.jwtServ
}

// GetTenantService returns an initialized tenant service object.
func (i *Initializer) GetTenantService() *service.TenantService {
	if i.tenServ == nil {
		i.tenServ = service.NewTenantService(
			i.conf,
			i.GetRedis().GetTenantRepo(),
			i.mq)
	}
	return i.tenServ
}

// GetRabbitMqService returns an initialized RabbitMQ service object.
func (i *Initializer) GetRabbitMqService() *service.RabbitMqService {
	if i.rabServ == nil {
		i.rabServ = service.NewRabbitMQService(
			i.conf,
			i.GetRedis().GetImageRepo())
	}
	return i.rabServ
}

// --- API controller functions ---

// GetImageApiController returns an initialized image API controller object.
func (i *Initializer) GetImageApiController() *ac.ImageController {
	if i.imgACtrl == nil {
		i.imgACtrl = ac.NewImageController(i.GetImageService(), i.GetTenantService())
	}
	return i.imgACtrl
}

// GetProbeApiController returns an initialized probe API controller object.
func (i *Initializer) GetProbeApiController() *ac.ProbeController {
	if i.prbACtrl == nil {
		i.prbACtrl = ac.NewProbeController()
	}
	return i.prbACtrl
}

// GetMetricsApiController returns an initialized metrics API controller object.
func (i *Initializer) GetMetricsApiController() *ac.MetricsController {
	if i.metACtrl == nil {
		i.metACtrl = ac.NewMetricsController()
	}
	return i.metACtrl
}

// --- API middleware functions ---

// GetTransactionApiMiddleware returns an initialized tx API middleware object.
func (i *Initializer) GetTransactionApiMiddleware() *am.TransactionMiddleware {
	if i.traAMdw == nil {
		i.traAMdw = am.NewTransactionMiddleware()
	}
	return i.traAMdw
}

// GetErrorApiMiddleware returns a initialized error API middleware object.
func (i *Initializer) GetErrorApiMiddleware() *am.ErrorMiddleware {
	if i.errAMdw == nil {
		i.errAMdw = am.NewErrorMiddleware()
	}
	return i.errAMdw
}

// GetSecurityApiMiddleware returns an initialized security API middleware object.
func (i *Initializer) GetSecurityApiMiddleware() *am.SecurityMiddleware {
	if i.secAMdw == nil {
		i.secAMdw = am.NewSecurityMiddleware()
	}
	return i.secAMdw
}

// GetAuthCheckApiMiddleware returns an initialized auth check API middleware object.
func (i *Initializer) GetAuthCheckApiMiddleware() *am.AuthCheckMiddleware {
	if i.authAMdw == nil {
		i.authAMdw = am.NewAuthCheckMiddleware(i.GetJWTService())
	}
	return i.authAMdw
}

// GetTokenCheckApiMiddleware returns an initialized token check API middleware object.
func (i *Initializer) GetTokenCheckApiMiddleware() *am.TokenCheckMiddleware {
	if i.tokenAMdw == nil {
		i.tokenAMdw = am.NewTokenCheckMiddleware()
	}
	return i.tokenAMdw
}
