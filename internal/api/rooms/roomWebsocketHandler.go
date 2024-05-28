package rooms_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// RoomWebsocketHandler implements HandlerInterface.
type RoomWebsocketHandler struct{}

func NewRoomWebsocketHandler() *RoomWebsocketHandler {
	return &RoomWebsocketHandler{}
}

func (*RoomWebsocketHandler) Pattern() string {
	return "/:code/ws"
}

func (*RoomWebsocketHandler) Handler(upgrader *websocket.Upgrader, log *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

		if nil != err {
			log.Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error connecting to websocket",
			})
		}

		defer func() {
			err := conn.Close()
			if nil != err {
				log.Error(err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error connecting to websocket",
				})
			}
		}()

		code := ctx.Param("code")
		if "" == code {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid code",
			})
		}

	}
}
