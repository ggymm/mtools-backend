package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var SetHandler = wire.NewSet(
	ConfigHandlerSet,
	DatabaseHandlerSet,
	CoderHandlerSet,
)

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return err
	}
	return nil
}

// ParseQuery 解析Query参数
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return err
	}
	return nil
}

// ParseForm 解析Form请求
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return err
	}
	return nil
}

// 返回业务JSON数据
func returnJson(c *gin.Context, success bool, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"data":    data,
	})
}

// 返回成功业务JSON数据
func returnSuccess(c *gin.Context, data interface{}) {
	returnJson(c, true, data)
}

func returnFailed(c *gin.Context, error interface{}) {
	returnJson(c, false, error)
}

func httpError(c *gin.Context, httpCode int, message interface{}) {
	c.JSON(httpCode, gin.H{"message": message})
}

func validatorErrorData(err error) string {
	var s string
	switch err.(type) {
	case validator.ValidationErrors:
		for _, err := range err.(validator.ValidationErrors) {
			if err != nil {
				s += "{" + err.Field() + "字段不符合规则}"
			}
		}
		return s
	case *json.SyntaxError:
		return err.Error()
	default:
		return err.Error()
	}
}
