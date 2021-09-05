package handler

import (
	"mtools-backend/model"
	"mtools-backend/schema"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var PostmanHandlerSet = wire.NewSet(wire.Struct(new(PostmanHandler), "*"))

type PostmanHandler struct {
	PostmanModel model.PostmanModel
}

func (h *PostmanHandler) GetTree(c *gin.Context) {
	if tree, err := h.PostmanModel.GetTree(); err != nil {
		returnFailed(c, err.Error())
		return
	} else {
		returnSuccess(c, tree)
		return
	}
}

func (h *PostmanHandler) Create(c *gin.Context) {
	collection := new(schema.CreateCollectionReq)
	err := ParseJSON(c, &collection)
	if err != nil {
		returnFailed(c, validatorErrorData(err))
		return
	}
	if err := h.PostmanModel.Create(collection); err != nil {
		returnFailed(c, err.Error())
		return
	} else {
		returnSuccess(c, nil)
		return
	}
}
