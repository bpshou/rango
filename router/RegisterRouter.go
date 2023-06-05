package router

import (
	"rango/app/controller/api"
	"rango/app/controller/use"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
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
	}
	routerUse := engine.Group("/use")
	{
		routerUse.GET("/mongo", use.Mongo{}.Mongo)
		routerUse.GET("/glog", use.Glog{}.Glog)
		routerUse.GET("/logrus", use.Logrus{}.Logrus)
	}
}
