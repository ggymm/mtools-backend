package main

import (
	"go.uber.org/zap"
	"net/http"

	"mtools-backend/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

type Router struct {
	Logger *zap.SugaredLogger
}

func (r *Router) NewRouter() (router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	router.Use(middleware.Cors())
	router.Use(middleware.ErrHandler(r.Logger))

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "start success")
	})
	v1 := router.Group("/api/v1/")
	{
		println(v1)
	}

	return
}
