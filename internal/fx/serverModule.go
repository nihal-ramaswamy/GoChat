package fx_utils

import (
	"context"
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	serverconfig "github.com/nihal-ramaswamy/GoVid/internal/config/server"
	"github.com/nihal-ramaswamy/GoVid/internal/dto"
	middleware "github.com/nihal-ramaswamy/GoVid/internal/middleware/log"
	"github.com/nihal-ramaswamy/GoVid/internal/routes"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func newServerEngine(
	lc fx.Lifecycle,
	rdb_auth *redis.Client,
	rdb_ws *redis.Client,
	config *serverconfig.Config,
	log *zap.Logger,
	upgrader *websocket.Upgrader,
	db *sql.DB,
	ctx context.Context,
	conferenceWsDto *dto.ConferenceWsDto,
) *gin.Engine {
	gin.SetMode(config.GinMode)

	server := gin.Default()
	server.Use(cors.New(config.Cors))
	server.Use(middleware.DefaultStructuredLogger(log))
	server.Use(gin.Recovery())

	routes.NewRoutes(server, upgrader, log, db, rdb_auth, ctx, rdb_ws, conferenceWsDto)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting server on port", zap.String("port", config.Port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping server")
			defer func() {
				err := log.Sync()
				if nil != err {
					log.Error(err.Error())
				}
			}()

			return nil
		},
	})

	return server
}

var serverModule = fx.Module(
	"serverModule",
	fx.Provide(
		fx.Annotate(
			newServerEngine,
			fx.ParamTags(``, `name:"auth_rdb"`, `name:"ws_rdb"`),
		),
	),
)
