package service

import (
	"pc-ziegert.de/media_service/common/config"
	ac "pc-ziegert.de/media_service/service/api/controller"
	am "pc-ziegert.de/media_service/service/api/middleware"
	"pc-ziegert.de/media_service/service/db"
	"pc-ziegert.de/media_service/service/mq"
	"pc-ziegert.de/media_service/service/service"
)

// Initializer initializes application components.
type Initializer struct {
	conf *config.Config

	mq    *mq.RabbitMQ
	redis *db.Redis

	jobServ    *service.JobService
	imageServ  *service.ImageService
	jwtServ    *service.JWTService
	tenantServ *service.TenantService

	imageACtrl *ac.ImageController

	traAMdw  *am.TransactionMiddleware
	errAMdw  *am.ErrorMiddleware
	secAMdw  *am.SecurityMiddleware
	authAMdw *am.AuthCheckMiddleware
}

// NewInitializer creates a new initializer.
func NewInitializer(conf *config.Config) *Initializer {
	return &Initializer{conf: conf}
}

// --- Config functions ---

// GetConfig returns a initialized config object.
func (i *Initializer) GetConfig() *config.Config {
	return i.conf
}

// --- RabbitMQ functions ---

// GetRabbitMQ returns a initialized RabbitMQ object.
func (i *Initializer) GetRabbitMQ() *mq.RabbitMQ {
	if i.mq == nil {
		i.mq = mq.NewRabbtmq(i.conf)
	}
	return i.mq
}

// --- Database functions ---

// GetRedis returns a initialized Redis object.
func (i *Initializer) GetRedis() *db.Redis {
	if i.redis == nil {
		i.redis = db.NewReddis(i.conf)
	}
	return i.redis
}

// --- Service functions ---

// GetJobService returns a initialized job service object
func (init *Initializer) GetJobService() *service.JobService {
	if init.jobServ == nil {
		init.jobServ = service.NewJobService()
	}
	return init.jobServ
}

// GetImageService returns a initialized image service object.
func (i *Initializer) GetImageService() *service.ImageService {
	if i.imageServ == nil {
		i.imageServ = service.NewImageService(
			i.conf,
			i.GetRedis().GetImageRepo(),
			i.mq)
	}
	return i.imageServ
}

// GetJWTService returns a initialized jwt service object.
func (i *Initializer) GetJWTService() *service.JWTService {
	if i.jwtServ == nil {
		i.jwtServ = service.NewJWTService(
			i.conf.Auth.Jwt.PublicKey)
	}
	return i.jwtServ
}

// GetTenantService returns a initialized tenant service object.
func (i *Initializer) GetTenantService() *service.TenantService {
	if i.tenantServ == nil {
		i.tenantServ = service.NewTenantService(
			i.conf,
			i.GetRedis().GetTenantRepo(),
			i.mq)
	}
	return i.tenantServ
}

// --- API controller functions ---

// GetImageApiController returns a initialized image API controller object.
func (i *Initializer) GetImageApiController() *ac.ImageController {
	if i.imageACtrl == nil {
		i.imageACtrl = ac.NewImageController(i.GetImageService(), i.GetTenantService())
	}
	return i.imageACtrl
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

// GetSecurityApiMiddleware returns a initialized security API middleware object.
func (init *Initializer) GetSecurityApiMiddleware() *am.SecurityMiddleware {
	if init.secAMdw == nil {
		init.secAMdw = am.NewSecurityMiddleware()
	}
	return init.secAMdw
}

// GetAuthCheckApiMiddleware returns a initialized auth check API middleware object.
func (i *Initializer) GetAuthCheckApiMiddleware() *am.AuthCheckMiddleware {
	if i.authAMdw == nil {
		i.authAMdw = am.NewAuthCheckMiddleware(i.GetJWTService())
	}
	return i.authAMdw
}
