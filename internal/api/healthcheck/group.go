package healthcheck_api

import "github.com/nihal-ramaswamy/GoVid/internal/interfaces"

type HealthCheckGroup struct {
	routeHandlers []interfaces.HandlerInterface
}

func (*HealthCheckGroup) Group() string {
	return "/healthcheck"
}

func (h *HealthCheckGroup) RouteHandlers() []interfaces.HandlerInterface {
	return h.routeHandlers
}

func NewHealthCheckGroup() *HealthCheckGroup {
	handlers := []interfaces.HandlerInterface{
		NewHealthCheckHandler(),
	}

	return &HealthCheckGroup{
		routeHandlers: handlers,
	}
}
