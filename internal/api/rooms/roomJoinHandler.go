package rooms_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"github.com/nihal-ramaswamy/GoVid/internal/utils"
)

// RoomJoinHandler implements HandlerInterface.
type RoomJoinHandler struct{}

func NewRoomJoinHandler() *RoomJoinHandler {
	return &RoomJoinHandler{}
}

func (RoomJoinHandler) Pattern() string {
	return "/:code"
}

func (*RoomJoinHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roomCode := ctx.Param("code")

		if "" == roomCode {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid code",
			})
		}

		utils.JoinRoom(roomCode)
	}
}

func (*RoomJoinHandler) RequestMethod() string {
	return constants.GET
}
