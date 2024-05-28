package fx_utils

import (
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
)

var WebsocketModule = fx.Module(
	"WebsocketModule",
	fx.Provide(provideUpgrader),
)

func provideUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}
