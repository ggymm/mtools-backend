package main

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

var AppSet = wire.NewSet(wire.Struct(new(App), "*"))

type App struct {
	Router *Router
	Logger *zap.SugaredLogger
}
