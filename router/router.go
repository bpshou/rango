package router

import (
	"rango/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()
	// 全局中间件注册
	engine.Use(middleware.Init, middleware.CrossSetHeader)
	// 注册路由
	ApiRouter(engine)
	UseRouter(engine)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	engine.Run("0.0.0.0:9090")
}
