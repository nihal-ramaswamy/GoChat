package chat_api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"github.com/nihal-ramaswamy/GoVid/internal/dto"
	"github.com/nihal-ramaswamy/GoVid/internal/utils"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ChatHandler struct {
	log             *zap.Logger
	upgrader        *websocket.Upgrader
	rdb             *redis.Client
	ctx             context.Context
	conferenceWsDto *dto.ConferenceWsDto

	middlewares []gin.HandlerFunc
}

func NewChatHandler(
	upgrader *websocket.Upgrader,
	log *zap.Logger,
	rdb_ws *redis.Client,
	ctx context.Context,
	conferenceWsDto *dto.ConferenceWsDto,
) *ChatHandler {
	return &ChatHandler{
		log:             log,
		upgrader:        upgrader,
		rdb:             rdb_ws,
		ctx:             ctx,
		conferenceWsDto: conferenceWsDto,
	}
}

func (*ChatHandler) Pattern() string {
	return "/:code"
}

func (*ChatHandler) RequestMethod() string {
	return constants.GET
}

func (c *ChatHandler) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *ChatHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// code := ctx.Param("code")
		// email := ctx.GetString("email")

		conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		utils.HandleErrorAndAbortWithError(ctx, err, c.log)

		defer func() {
			err := conn.Close()
			utils.HandleErrorAndAbortWithError(ctx, err, c.log)
		}()

		c.conferenceWsDto.AddConnection(conn)
		c.conferenceWsDto.HandleWs(conn)
	}
}
