package api

import (
	"rango/app/controller"
	"rango/app/service"

	"github.com/gin-gonic/gin"
)

type Viper struct {
	controller.Base
}

// 控制器
func (the Viper) ViperConfig(c *gin.Context) {
	// 服务
	service.ViperConfig(c)
}
