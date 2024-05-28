package routes

import (
	"github.com/gin-gonic/gin"
	healthcheck_api "github.com/nihal-ramaswamy/GoVid/internal/api/healthcheck"
	rooms_api "github.com/nihal-ramaswamy/GoVid/internal/api/rooms"
)

func NewRoutes(server *gin.Engine) {
	healthCheckHandler := healthcheck_api.NewHealthCheckHandler()
	roomCreateHandler := rooms_api.NewRoomCretateHandler()

	healthcheck := server.Group("/healthcheck")
	{
		healthcheck.GET(healthCheckHandler.Pattern(), healthCheckHandler.Handler())
	}

	room := server.Group("/room")
	{
		room.GET(roomCreateHandler.Pattern(), roomCreateHandler.Handler())
	}
}
