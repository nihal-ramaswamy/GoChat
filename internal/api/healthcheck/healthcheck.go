package healthcheck_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Any API in this call needs to be prefixed with /room.

// HealthCheckHandler implements HandlerInterface.
type HealthCheckHandler struct {
}

func (*HealthCheckHandler) Pattern() string {
	return "/healthcheck"
}

func (*HealthCheckHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}
