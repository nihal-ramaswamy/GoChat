package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	healthcheck_api "github.com/nihal-ramaswamy/GoVid/internal/api/healthcheck"
	rooms_api "github.com/nihal-ramaswamy/GoVid/internal/api/rooms"
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"github.com/nihal-ramaswamy/GoVid/internal/interfaces"
	"go.uber.org/zap"
)

func NewRoutes(
	server *gin.Engine,
	upgrader *websocket.Upgrader,
	log *zap.Logger,
) {
	serverGroupHandlers := []interfaces.ServerGroupInterface{
		healthcheck_api.NewHealthCheckGroup(),
		rooms_api.NewRoomGroup(log, upgrader),
	}

	for _, serverGroupHandler := range serverGroupHandlers {
		newGroup(server, serverGroupHandler)
	}
}

func newGroup(server *gin.Engine, groupHandler interfaces.ServerGroupInterface) {
	group := server.Group(groupHandler.Group())
	{
		for _, route := range groupHandler.RouteHandlers() {
			newRoute(group, route)
		}
	}
}

func newRoute(server *gin.RouterGroup, routeHandler interfaces.HandlerInterface) {
	switch routeHandler.RequestMethod() {
	case constants.GET:
		server.GET(routeHandler.Pattern(), routeHandler.Handler())
	case constants.POST:
		server.POST(routeHandler.Pattern(), routeHandler.Handler())
	}
}
