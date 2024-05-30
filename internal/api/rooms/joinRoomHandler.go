package rooms_api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"go.uber.org/zap"
)

type JoinRoomHandler struct {
	db          *sql.DB
	log         *zap.Logger
	middlewares []gin.HandlerFunc
}

func NewJoinRoomHandler(db *sql.DB, log *zap.Logger) *JoinRoomHandler {
	return &JoinRoomHandler{
		db:          db,
		log:         log,
		middlewares: []gin.HandlerFunc{},
	}
}

func (*JoinRoomHandler) Pattern() string {
	return "/:code"
}

func (*JoinRoomHandler) RequestMethod() string {
	return constants.POST
}

func (c *JoinRoomHandler) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

// func (c *JoinRoomHandler) Handler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		email := ctx.GetString("email")
// 		code := ctx.Param("code")
//
// 	}
// }
