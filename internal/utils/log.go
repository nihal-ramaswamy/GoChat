package utils

import (
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"go.uber.org/zap"
)

func NewProduction() *zap.Logger {
	env := GetDotEnvVariable(constants.ENV)

	switch env {
	case "release":
		return zap.Must(zap.NewProduction())
	case "debug":
		return zap.Must(zap.NewDevelopment())
	default:
		return zap.Must(zap.NewDevelopment())
	}
}
