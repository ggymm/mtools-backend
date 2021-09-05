// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"mtools-backend/config"
	"mtools-backend/database"
	"mtools-backend/handler"
	"mtools-backend/logger"
	"mtools-backend/model"

	"github.com/google/wire"
)

// BuildApp 生成注入器
func BuildApp() (*App, func(), error) {
	wire.Build(
		config.InitConfig,
		logger.InitLogger,
		database.InitXormDB,
		handler.SetHandler,
		model.SetModel,
		RouterSet,
		AppSet,
	)
	return new(App), nil, nil
}
