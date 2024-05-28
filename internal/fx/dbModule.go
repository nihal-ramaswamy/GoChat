package fx_utils

import (
	"context"

	"github.com/nihal-ramaswamy/GoVid/internal/db"
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
