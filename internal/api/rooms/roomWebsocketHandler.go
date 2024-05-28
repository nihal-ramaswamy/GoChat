package rooms_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"go.uber.org/zap"
)

// RoomWebsocketHandler implements HandlerInterface.
type RoomWebsocketHandler struct {
	log      *zap.Logger
	upgrader *websocket.Upgrader
}

func NewRoomWebsocketHandler(log *zap.Logger, upgrader *websocket.Upgrader) *RoomWebsocketHandler {
	return &RoomWebsocketHandler{
		log:      log,
		upgrader: upgrader,
	}
}

func (*RoomWebsocketHandler) Pattern() string {
	return "/:code/ws"
}

func (*RoomWebsocketHandler) RequestMethod() string {
	return constants.GET
}

func (r *RoomWebsocketHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conn, err := r.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		r.log.Info("", zap.String("conn", conn.Subprotocol()))

		if nil != err {
			r.log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Error connecting to websocket",
			})
		}

		defer conn.Close()

		code := ctx.Param("code")
		if "" == code {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid code",
			})
		}

		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				r.log.Error("read:", zap.String("error", err.Error()))
				break
			}

			r.log.Info("Message received", zap.String("received", string(message)), zap.String("type", string(websocket.FormatCloseMessage(mt, string(message)))))
			err = conn.WriteMessage(mt, message)
			if err != nil {
				r.log.Error("read:", zap.String("error", err.Error()))
				break
			}
		}

	}
}
