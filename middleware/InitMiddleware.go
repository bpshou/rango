package middleware

import (
	"github.com/gin-gonic/gin"
)

// 初始化配置
func Init(c *gin.Context) {
	// 校验权限

	c.Next()
}
