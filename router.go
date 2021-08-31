package main

import (
	"go.uber.org/zap"
	"mtools-backend/handler"
	"net/http"

	"mtools-backend/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

type Router struct {
	Logger          *zap.SugaredLogger
	ConfigHandler   *handler.ConfigHandler
	DatabaseHandler *handler.DatabaseHandler
	CoderHandler    *handler.CoderHandler
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
		// 配置文件功能
		config := v1.Group("config/").Use(middleware.CheckUser())
		{
			config.GET("/", r.ConfigHandler.Get)
		}

		// 数据库功能
		database := v1.Group("database/").Use(middleware.CheckUser())
		{
			database.GET("get-db-list", r.DatabaseHandler.GetTableList)
			database.GET("get-table-list", r.DatabaseHandler.GetTableList)
		}

		// 代码生成器功能
		coder := v1.Group("coder/").Use(middleware.CheckUser())
		{
			coder.POST("gen-code", r.CoderHandler.GenCode)
		}
	}

	return
}
