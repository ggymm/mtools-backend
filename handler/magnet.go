package handler

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"mtools-backend/config"
)

var MagnetHandlerSet = wire.NewSet(wire.Struct(new(MagnetHandler), "*"))

type MagnetHandler struct {
	Logger *zap.SugaredLogger
	Config *config.GlobalConfig
}
