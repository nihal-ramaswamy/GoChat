package auth_api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type LogoutUserHandler struct {
	ctx context.Context
	rdb *redis.Client
	log *zap.Logger
}

func NewLogoutUserHandler(rdb *redis.Client, ctx context.Context, log *zap.Logger) *LogoutUserHandler {
	return &LogoutUserHandler{
		rdb: rdb,
		ctx: ctx,
		log: log,
	}
}

func (*LogoutUserHandler) Pattern() string {
	return "/signout"
}

func (*LogoutUserHandler) RequestMethod() string {
	return constants.POST
}

func (l *LogoutUserHandler) Handler() gin.HandlerFunc {

	return func(c *gin.Context) {
		email := c.GetString("email")
		_, err := l.rdb.Del(l.ctx, email).Result()

		if err != nil {
			l.log.Error("Error deleting token from rdb")
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "ok"})
	}
}
