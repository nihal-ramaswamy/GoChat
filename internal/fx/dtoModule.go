package fx_utils

import (
	"github.com/nihal-ramaswamy/GoVid/internal/dto"
	"go.uber.org/fx"
)

var DtoModule = fx.Module(
	"DtoModule",
	fx.Provide(dto.NewRoom),
)
