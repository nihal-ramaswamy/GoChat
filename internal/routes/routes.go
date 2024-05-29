package routes

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	auth_api "github.com/nihal-ramaswamy/GoVid/internal/api/auth"
	healthcheck_api "github.com/nihal-ramaswamy/GoVid/internal/api/healthcheck"
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"github.com/nihal-ramaswamy/GoVid/internal/interfaces"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func NewRoutes(
	server *gin.Engine,
	upgrader *websocket.Upgrader,
	log *zap.Logger,
	db *sql.DB,
	rdb *redis.Client,
	ctx context.Context,
) {
	serverGroupHandlers := []interfaces.ServerGroupInterface{
		healthcheck_api.NewHealthCheckGroup(),
		auth_api.NewAuthGroup(db, rdb, ctx, log),
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
