package router

import (
	"rango/app/controller/api"

	"github.com/gin-gonic/gin"
)

func ApiRouter(engine *gin.Engine) {
	router := engine.Group("/api")
	{
		router.GET("/", api.Index{}.Index)
		router.GET("/ping", api.Index{}.Ping)
		router.GET("/mysql", api.Index{}.Mysql)
		router.GET("/viper/config", api.Viper{}.ViperConfig)
		router.POST("/cmd", api.Command{}.Index)
		router.GET("/aes", api.Aes{}.Index)
	}
}
