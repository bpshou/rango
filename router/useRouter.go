package router

import (
	"rango/app/controller/use"

	"github.com/gin-gonic/gin"
)

func UseRouter(engine *gin.Engine) {
	router := engine.Group("/use")
	{
		router.GET("/mongo", use.Mongo{}.Mongo)
		router.GET("/glog", use.Glog{}.Glog)
		router.GET("/logrus", use.Logrus{}.Logrus)
	}
}
