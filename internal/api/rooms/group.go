package rooms_api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/GoVid/internal/interfaces"
	"go.uber.org/zap"
)

type RoomsApi struct {
	routeHandlers []interfaces.HandlerInterface
	middlewares   []gin.HandlerFunc
}

func (*RoomsApi) Group() string {
	return "/rooms"
}

func (h *RoomsApi) RouteHandlers() []interfaces.HandlerInterface {
	return h.routeHandlers
}

func NewRoomsApi(db *sql.DB, log *zap.Logger) *RoomsApi {
	handlers := []interfaces.HandlerInterface{
		NewCreateRoomHandler(db, log),
	}

	return &RoomsApi{
		routeHandlers: handlers,
		middlewares:   []gin.HandlerFunc{},
	}
}

func (*RoomsApi) AuthRequired() bool {
	return false
}

func (h *RoomsApi) Middlewares() []gin.HandlerFunc {
	return h.middlewares
}
