package interfaces

type ServerGroupInterface interface {
	Group() string
	RouteHandlers() []HandlerInterface
}
