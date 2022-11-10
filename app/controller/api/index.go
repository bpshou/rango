package api

import (
	"net/http"
	"rango/app/controller"
	"rango/app/service"

	"github.com/gin-gonic/gin"
)

type Index struct {
	controller.Base
}

type Params struct {
	Name string `form:"name" json:"name" xml:"name"`
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

func (this Index) Params(c *gin.Context) {
	// 获取json参数
	var params Params
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    200,
			"message": err.Error(),
		})
		return
	}

	// Query 用来接收url上的参数，例如 http://localhost/api/params?name=xiaoming
	var name = c.Query("name")

	// Param 用来接收路由上的参数，例如 http://localhost/api/params/18?name=xiaoming
	var id = c.Param("id")

	c.JSON(200, gin.H{
		"code":   200,
		"params": params,
		"query":  name,
		"id":     id,
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
