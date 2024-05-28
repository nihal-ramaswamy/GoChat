package rooms_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"github.com/nihal-ramaswamy/GoVid/internal/utils"
)

// Any API in this call needs to be prefixed with /room

// RoomCreateHandler implements HandlerInterface.
type RoomCreateHandler struct {
}

func (*RoomCreateHandler) Pattern() string {
	return "/create"
}

func (*RoomCreateHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuidString := utils.NewUUID(constants.UUID_LENGTH)
		ctx.Redirect(http.StatusTemporaryRedirect, "/room/"+uuidString)

	}
}

func NewRoomCretateHandler() *RoomCreateHandler {
	return &RoomCreateHandler{}
}

func (*RoomCreateHandler) RequestMethod() string {
	return constants.GET
}
