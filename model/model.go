package model

import "github.com/google/wire"

var SetModel = wire.NewSet(
	DatabaseModelSet,
	PostmanModelSet,
)
