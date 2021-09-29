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
	DatabaseHandler *handler.DatabaseHandler
	CoderHandler    *handler.CoderHandler
	PostmanHandler  *handler.PostmanHandler
}

func (r *Router) NewRouter() (router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	router.Use(middleware.Cors())
	router.Use(middleware.ErrHandler(r.Logger))

	// router.StaticFS("/drawio", http.Dir("drawio"))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "start success")
	})
	v1 := router.Group("/api/v1/")
	{
		// 数据库
		database := v1.Group("database/").Use(middleware.CheckUser())
		{
			database.GET("list", r.DatabaseHandler.GetDatabaseList)
			database.GET("tableList", r.DatabaseHandler.GetTableList)
		}

		// 代码生成器
		coder := v1.Group("coder/").Use(middleware.CheckUser())
		{
			coder.POST("generateJavaCode", r.CoderHandler.GenJavaCode)
			coder.POST("generateGoCode", r.CoderHandler.GenGoCode)
		}

		// 接口调试工具
		postman := v1.Group("postman/")
		{
			collection := postman.Group("collection/").Use(middleware.CheckUser())
			{
				collection.GET("tree", r.PostmanHandler.GetTree)
				collection.POST("create", r.PostmanHandler.Create)
			}

		}

		// oss管理

		// 磁力链接搜索
		magnet := v1.Group("magnet/").Use(middleware.CheckUser())
		{
			magnet.GET("search", r.CoderHandler.GenJavaCode)
		}
	}

	return
}
