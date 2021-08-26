package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ConfigHandlerSet = wire.NewSet(wire.Struct(new(ConfigHandler), "*"))

type ConfigHandler struct {
}

func (h *ConfigHandler) Get(c *gin.Context) {
	returnSuccess(c, nil)
	return
}
