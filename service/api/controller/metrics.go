package controller

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// MetricsController handles requests for probe endpoints.
type MetricsController struct {
	p *ginprometheus.Prometheus
}

// NewMetricsController NewProbeController create a new metrics controller.
func NewMetricsController() *MetricsController {
	return &MetricsController{
		p: ginprometheus.NewPrometheus("gin"),
	}
}

func (metCtl *MetricsController) MetricsHandler(e *gin.Engine) {
	metCtl.p.Use(e)
}
