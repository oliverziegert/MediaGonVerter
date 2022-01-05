package main

import (
	ac "pc-ziegert.de/media_service/api/controller"
	am "pc-ziegert.de/media_service/api/middleware"
	"pc-ziegert.de/media_service/config"
	"pc-ziegert.de/media_service/service"
)

// Initializer initializes application components.
type Initializer struct {
	conf *config.Config

	jobServ   *service.JobService
	imageServ *service.ImageService

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
		i.imageServ = service.NewImageService(i.conf)
	}
	return i.imageServ
}

// --- API controller functions ---

// GetImageApiController returns a initialized image API controller object.
func (i *Initializer) GetImageApiController() *ac.ImageController {
	if i.imageACtrl == nil {
		i.imageACtrl = ac.NewImageController(i.GetImageService())
	}
	return i.imageACtrl
}

// --- API middleware functions ---

// GetTransactionApiMiddleware returns an initialized transaction API middleware object.
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
		i.authAMdw = am.NewAuthCheckMiddleware()
	}
	return i.authAMdw
}
