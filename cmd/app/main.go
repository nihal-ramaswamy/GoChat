package main

import (
	"github.com/gin-gonic/gin"
	serverconfig "github.com/nihal-ramaswamy/GoVid/internal/config/server"
	fx_utils "github.com/nihal-ramaswamy/GoVid/internal/fx"
	"github.com/nihal-ramaswamy/GoVid/internal/utils"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(utils.NewProduction),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),

		fx_utils.ConfigModule,
		fx_utils.MicroServicesModule,

		fx.Invoke(Invoke),
	).Run()
}

func Invoke(server *gin.Engine, config *serverconfig.Config) {
	server.Run(config.Port)
}
