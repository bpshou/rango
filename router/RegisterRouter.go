package router

import (
	"rango/app/controller/api"
	"rango/app/controller/use"

	docs "rango/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRouter(engine *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"
	router := engine.Group("/api")
	{
		router.GET("/", api.Index{}.Index)
		router.GET("/totp", api.Totp{}.Secret)
		router.GET("/params/:id", api.Index{}.Params)
		router.GET("/ping", api.Index{}.Ping)
		router.GET("/mysql", api.Index{}.Mysql)
		router.GET("/viper/config", api.Viper{}.ViperConfig)
		router.POST("/cmd", api.Command{}.Index)
		router.GET("/aes", api.Aes{}.Index)
		router.GET("/kafka", api.Kafka{}.Start)
		router.GET("/jwt", api.Jwt{}.Index)
		router.GET("/secret/add", api.Secret{}.Create)
		router.GET("/secret/select", api.Secret{}.Decrypt)
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	routerUse := engine.Group("/use")
	{
		routerUse.GET("/mongo", use.Mongo{}.Mongo)
		routerUse.GET("/glog", use.Glog{}.Glog)
		routerUse.GET("/logrus", use.Logrus{}.Logrus)
	}
}
