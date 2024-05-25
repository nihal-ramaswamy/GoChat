package main

import (
	serverconfig "example.go_fx_gin/internal/config/server"
	fx_utils "example.go_fx_gin/internal/fx"
	"example.go_fx_gin/internal/utils"
	"github.com/gin-gonic/gin"
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
