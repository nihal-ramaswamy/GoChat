package healthcheck_api

import "github.com/gin-gonic/gin"

type HealthCheckHandler struct {
}

func (*HealthCheckHandler) Pattern() string {
	return "/healthcheck"
}

func (*HealthCheckHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	}
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}
