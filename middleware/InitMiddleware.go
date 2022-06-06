package middleware

import (
	"rango/tools/logger"
	"rango/tools/viper"
	"rango/tools/xorm"

	"github.com/gin-gonic/gin"
)

// 初始化配置
func Init(c *gin.Context) {
	viper.ImportConfig()
	logger.Init()
	xorm.EngineGroup()
	c.Next()
}
