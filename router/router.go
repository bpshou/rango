package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"rango/app/service"
	log "github.com/sirupsen/logrus"
)

func Run() {
	route := gin.Default()
	route.Use(CrossSetHeader())
	route.GET("/", service.Index)
	route.GET("/ping", service.Ping)
	route.GET("/http", service.Http)
	route.GET("/mongo", service.Mongo)
	route.GET("/aes", service.Aes)
	route.GET("/use/glog", service.UseGlog)
	route.GET("/use/logrus", service.UseLogrus)
	route.POST("/cmd", service.Command)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	route.Run("0.0.0.0:9090")
}

// 设置跨域
func CrossSetHeader() gin.HandlerFunc {
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
