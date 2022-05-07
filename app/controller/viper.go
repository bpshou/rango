package controller

import (
    "rango/app/service"
    "github.com/gin-gonic/gin"
)

// 控制器
func ViperConfig(c *gin.Context)  {
    // 服务
    service.ViperConfig(c)
}