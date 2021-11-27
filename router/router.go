package router

import (
	"github.com/gin-gonic/gin"
	"rango/app/Cross"
	"rango/service"
)

func Run() {
	route := gin.Default()
	route.Use(Cross.SetHeader())
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
