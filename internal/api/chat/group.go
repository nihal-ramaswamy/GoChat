package chat_api

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nihal-ramaswamy/GoVid/internal/dto"
	"github.com/nihal-ramaswamy/GoVid/internal/interfaces"
	auth_middleware "github.com/nihal-ramaswamy/GoVid/internal/middleware/auth"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ChatApi struct {
	middlewares   []gin.HandlerFunc
	routeHandlers []interfaces.HandlerInterface
}

func NewChatApi(
	pdb *sql.DB,
	rdb_auth *redis.Client,
	ctx context.Context,
	log *zap.Logger,
	upgrader *websocket.Upgrader,
	rdb_ws *redis.Client,
	roomDto *dto.Room,
) *ChatApi {
	handlers := []interfaces.HandlerInterface{
		NewChatHandler(upgrader, log, rdb_ws, ctx, roomDto),
	}

	return &ChatApi{
		middlewares:   []gin.HandlerFunc{auth_middleware.AuthMiddleware(pdb, rdb_auth, ctx, log)},
		routeHandlers: handlers,
	}
}

func (*ChatApi) Group() string {
	return "/chat"
}

func (c *ChatApi) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *ChatApi) RouteHandlers() []interfaces.HandlerInterface {
	return c.routeHandlers
}
