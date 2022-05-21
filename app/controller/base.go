package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Base struct {
}

// 请求成功
func (this Base) Success(code int, params ...interface{}) {
	var c *gin.Context
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": params,
	})
	return
}

// 请求错误
func (this Base) Error(code int, params ...interface{}) {
	var c *gin.Context
	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"data": params,
	})
	return
}
