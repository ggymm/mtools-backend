package main

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"mtools-backend/config"
)

var AppSet = wire.NewSet(wire.Struct(new(App), "*"))

type App struct {
	Router *Router
	Config *config.GlobalConfig
	Logger *zap.SugaredLogger
}
