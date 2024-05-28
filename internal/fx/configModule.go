package fx_utils

import (
	db_config "github.com/nihal-ramaswamy/GoVid/internal/config/db"
	serverconfig "github.com/nihal-ramaswamy/GoVid/internal/config/server"
	"go.uber.org/fx"
)

var ConfigModule = fx.Module(
	"Config",
	fx.Provide(serverconfig.Default),
	fx.Provide(db_config.GetPsqlInfoDefault),
	fx.Provide(db_config.DefaultRedisConfig),
)
