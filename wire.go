// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"mtools-backend/config"
	"mtools-backend/logger"

	"github.com/google/wire"
)

// BuildApp 生成注入器
func BuildApp() (*App, func(), error) {
	wire.Build(
		logger.InitLogger,
		config.InitConfig,
		RouterSet,
		AppSet,
	)
	return new(App), nil, nil
}
