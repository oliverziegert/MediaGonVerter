package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ProbeController handles requests for probe endpoints.
type ProbeController struct {
}

// NewProbeController create a new user controller.
func NewProbeController() *ProbeController {
	return &ProbeController{}
}

func (prbCtrl *ProbeController) GetHealthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		return
	}
}

func (prbCtrl *ProbeController) GetReadyHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		return
	}
}

func (prbCtrl *ProbeController) GetStartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		return
	}
}
