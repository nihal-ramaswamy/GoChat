package utils

import (
	healthcheck_api "example.go_fx_gin/internal/api/healthcheck"
	"github.com/gin-gonic/gin"
)

func NewRoutes(server *gin.Engine) {
	healthCheckHandler := healthcheck_api.NewHealthCheckHandler()
	healthcheck := server.Group("/healthcheck")
	{
		healthcheck.GET(healthCheckHandler.Pattern(), healthCheckHandler.Handler())
	}
}
