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

// @Tags 		Api模块
// @Summary 	首页index
// @Description 首页index入口，默认返回为一个index字符串
// @Accept 		json
// @Produce 	json
// @Success 	200 {string} index
// @Router 		/api/ [get]
func (the Index) Index(c *gin.Context) {
	the.Success(c, 200, "index")
}

func (the Index) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "pong",
	})
}

// @Tags 		Api模块
// @Summary   	测试请求参数
// @Accept 		json
// @Produce 	json
// @Param     	data   body      Params			true  "name入参，测试入参"
// @Success   	200    {map}  gin.H{code=string}		"响应的body体"
// @Router    	/api/params/:id  [get]
func (the Index) Params(c *gin.Context) {
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

func (the Index) Mysql(c *gin.Context) {

	userInsert := service.Insert()
	userSelect := service.Select()
	userUpdate := service.Update()
	userDelete := service.Delete()
	userGetOne := service.GetOne()
	userGetCount := service.GetCount()

	c.JSON(200, gin.H{
		"code":         200,
		"message":      "mysql",
		"userInsert":   userInsert,
		"userSelect":   userSelect,
		"userUpdate":   userUpdate,
		"userDelete":   userDelete,
		"userGetOne":   userGetOne,
		"userGetCount": userGetCount,
	})
}
