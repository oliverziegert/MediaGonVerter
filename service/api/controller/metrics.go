package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	e.Use(metCtl.p.HandlerFunc())
	h := promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer, promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{DisableCompression: true}),
	)
	e.GET(metCtl.p.MetricsPath, func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})
}
