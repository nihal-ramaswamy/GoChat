package fx_utils

import (
	db_config "example.go_fx_gin/internal/config/db"
	serverconfig "example.go_fx_gin/internal/config/server"
	"go.uber.org/fx"
)

var ConfigModule = fx.Module(
	"Config",
	fx.Provide(serverconfig.Default),
	fx.Provide(db_config.GetPsqlInfoDefault),
	fx.Provide(db_config.DefaultRedisConfig),
)
