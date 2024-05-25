package fx_utils

import (
	"context"

	"example.go_fx_gin/internal/db"
	"go.uber.org/fx"
)

var dbModule = fx.Module(
	"DatabaseServices",
	fx.Provide(func() context.Context {
		return context.Background()
	}),

	fx.Provide(db.GetPostgresDbInstanceWithConfig),
	fx.Provide(db.GetRedisDbInstanceWithConfig),
)
