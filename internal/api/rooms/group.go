package rooms_api

import (
	"github.com/gorilla/websocket"
	"github.com/nihal-ramaswamy/GoVid/internal/interfaces"
	"go.uber.org/zap"
)

type RoomGroup struct {
	routeHandlers []interfaces.HandlerInterface
}

func (*RoomGroup) Group() string {
	return "/room"
}

func (r *RoomGroup) RouteHandlers() []interfaces.HandlerInterface {
	return r.routeHandlers
}

func NewRoomGroup(log *zap.Logger, upgrader *websocket.Upgrader) *RoomGroup {
	handlers := []interfaces.HandlerInterface{
		NewRoomCretateHandler(),
		NewRoomJoinHandler(),
		NewRoomWebsocketHandler(log, upgrader),
	}

	return &RoomGroup{
		routeHandlers: handlers,
	}
}
