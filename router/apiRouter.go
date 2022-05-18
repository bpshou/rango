package router

import (
	"rango/app/controller"
	"rango/app/service"

	"github.com/gin-gonic/gin"
)

func ApiRouter(engine *gin.Engine) {
	router := engine.Group("/api")
	{
		router.GET("/", service.Index)
		router.GET("/ping", service.Ping)
		router.GET("/viper/config", controller.ViperConfig)
	}
}
