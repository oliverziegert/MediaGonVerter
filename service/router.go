package service

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
	"pc-ziegert.de/media_service/common/constant"
	"time"
)

func ConfigureRouter(i *Initializer, router *gin.Engine) {
	if i.GetConfig().HTTP.Gzip.Enabled {
		router.Use(gzip.Gzip(gzip.BestCompression))
	}
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = i.GetConfig().HTTP.Origin.All.Enabled
	//corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(corsConfig))

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			param.Latency = param.Latency - param.Latency%time.Second
		}
		// 2022/03/27 15:39:26 [INFO ]: Using config file: /Users/oliver.ziegert/.pc-ziegert/media-service.yaml
		return fmt.Sprintf("%v [INFO ]: |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
}

func ConfigureApiRouting(i *Initializer, router *gin.Engine) {
	// Create API sub route
	ir := router.Group(fmt.Sprintf("%s/image", constant.ApiPath))
	nr := router.Group(fmt.Sprintf("%s/api/v1/nodes", constant.ApiPath))

	// Create protected middleware route
	ir.Use(i.GetErrorApiMiddleware().ServeHTTP())
	ir.Use(i.GetSecurityApiMiddleware().ServeHTTP())
	ir.Use(i.GetTokenCheckApiMiddleware().ServeHTTP())

	// Create protected middleware route
	nr.Use(i.GetErrorApiMiddleware().ServeHTTP())
	nr.Use(i.GetSecurityApiMiddleware().ServeHTTP())
	// use jwt for "/nodes" endpoints
	nr.Use(i.GetAuthCheckApiMiddleware().ServeHTTP())

	// Get controllers
	imgCtrl := i.GetImageApiController()

	// Add protected endpoints
	ir.HEAD("/:token/:size", imgCtrl.HeadImageHandler())
	ir.GET("/:token/:size", imgCtrl.GetImageHandler())

	nr.HEAD("/:nodeId/preview/:size", imgCtrl.HeadNodePreviewHandler())
	nr.GET("/:nodeId/preview/:size", imgCtrl.GetNodePreviewHandler())
}

func AddSwaggerUiHandlers(router *gin.Engine) {
	router.GET(constant.ApiPath, func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%s/swagger-ui/index.html", constant.ApiPath))
	})
	router.StaticFS(fmt.Sprintf("%s/swagger-ui/", constant.ApiPath), http.Dir("./service/swagger-ui"))
}

func AddRootHandlers(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, constant.ApiPath)
	})
}

func AddProbeHandlers(i *Initializer, router *gin.Engine) {
	// Get controllers
	prbCtrl := i.GetProbeApiController()

	router.GET("/healthz", prbCtrl.GetHealthHandler())
	router.GET("/readyz", prbCtrl.GetReadyHandler())
	router.GET("/startz", prbCtrl.GetStartHandler())
}

func AddMetricsHandlers(i *Initializer, router *gin.Engine) {
	// Get controllers
	metCtrl := i.GetMetricsApiController()

	metCtrl.MetricsHandler(router)
}
