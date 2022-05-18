package router

import (
	"rango/app/service"

	"github.com/gin-gonic/gin"
)

func UseRouter(engine *gin.Engine) {
	router := engine.Group("/use")
	{
		router.GET("/mongo", service.Mongo)
		router.GET("/aes", service.Aes)
		router.GET("/glog", service.UseGlog)
		router.GET("/logrus", service.UseLogrus)
		router.POST("/cmd", service.Command)
	}
}
