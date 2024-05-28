package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	healthcheck_api "github.com/nihal-ramaswamy/GoVid/internal/api/healthcheck"
	rooms_api "github.com/nihal-ramaswamy/GoVid/internal/api/rooms"
	"go.uber.org/zap"
)

func NewRoutes(
	server *gin.Engine,
	upgrader *websocket.Upgrader,
	log *zap.Logger,
) {
	// healthcheck
	healthCheckHandler := healthcheck_api.NewHealthCheckHandler()

	healthcheck := server.Group("/healthcheck")
	{
		healthcheck.GET(healthCheckHandler.Pattern(), healthCheckHandler.Handler())
	}

	// room
	roomCreateHandler := rooms_api.NewRoomCretateHandler()
	roomJoinHandler := rooms_api.NewRoomJoinHandler()
	roomWebSocketHandler := rooms_api.NewRoomWebsocketHandler(log, upgrader)

	room := server.Group("/room")
	{
		room.GET(roomCreateHandler.Pattern(), roomCreateHandler.Handler())
		room.GET(roomJoinHandler.Pattern(), roomJoinHandler.Handler())
		room.GET(roomWebSocketHandler.Pattern(), roomWebSocketHandler.Handler())
	}
}
