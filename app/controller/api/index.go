package api

import (
	"rango/app/controller"
	"rango/app/service"

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

func (this Index) Mysql(c *gin.Context) {

	insert := service.Insert()
	Select := service.Select()
	update := service.Update()
	delete := service.Delete()

	c.JSON(200, gin.H{
		"code":    200,
		"message": "mysql",
		"insert":  insert,
		"Select":  Select,
		"update":  update,
		"delete":  delete,
	})
}
