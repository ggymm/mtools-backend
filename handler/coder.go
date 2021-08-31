package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	"mtools-backend/config"
)

var CoderHandlerSet = wire.NewSet(wire.Struct(new(CoderHandler), "*"))

type CoderHandler struct {
	Logger *zap.SugaredLogger
	Config *config.GlobalConfig
}

func (h *CoderHandler) GenCode(c *gin.Context) {
	returnSuccess(c, nil)
	return
}
