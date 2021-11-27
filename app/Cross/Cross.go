package Cross

import (
	"net/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func SetHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		log.Debug(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE") // 服务器支持的所有跨域请求的方式
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
