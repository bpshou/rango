package api

import (
	"rango/app/controller"

	"github.com/gin-gonic/gin"
)

type Index struct {
	controller.Base
}

func (this Index) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "index",
	})
}

func (this Index) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "pong",
	})
}
