package api

import (
	"rango/app/controller"
	"rango/app/service"

	"github.com/gin-gonic/gin"
)

type Command struct {
	controller.Base
}

// 控制器
func (the Command) Index(c *gin.Context) {
	// 服务
	service.Command(c)
}
