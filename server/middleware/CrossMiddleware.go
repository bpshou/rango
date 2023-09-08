package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 设置跨域
func CrossSetHeader(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:90")
	// 服务器支持跨域请求的方式
	c.Header("Access-Control-Allow-Methods", "HEAD, POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	// 客户端请求服务端所能携带header的名单
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 客户端能拿到服务端响应header的名单
	c.Header("Access-Control-Expose-Headers", "")
	c.Header("Access-Control-Allow-Credentials", "true")
	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}
