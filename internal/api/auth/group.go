package auth_api

import (
	"context"
	"database/sql"

	"github.com/nihal-ramaswamy/GoVid/internal/interfaces"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type AuthGroup struct {
	routeHandlers []interfaces.HandlerInterface
}

func (*AuthGroup) Group() string {
	return "/auth"
}

func (h *AuthGroup) RouteHandlers() []interfaces.HandlerInterface {
	return h.routeHandlers
}

func NewAuthGroup(db *sql.DB, rdb *redis.Client, ctx context.Context, log *zap.Logger) *AuthGroup {
	handlers := []interfaces.HandlerInterface{
		NewNewUserHandler(db, log),
		NewLoginUserHandler(db, rdb, ctx, log),
		NewLogoutUserHandler(rdb, ctx, log),
	}

	return &AuthGroup{
		routeHandlers: handlers,
	}
}
