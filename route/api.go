package route

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"rango/service"
)

func Run() {
	route := gin.Default()

	route.GET("/", service.Index)
	route.GET("/ping", service.Ping)
	route.GET("/http", service.Http)
	route.GET("/mongo", service.Mongo)
	route.GET("/aes", service.Aes)

	fmt.Println("0.0.0.0:9090 Service start")
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	route.Run("0.0.0.0:9090")
}
