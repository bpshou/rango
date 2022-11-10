package router

import (
	"rango/middleware"
	"rango/tools/logger"
	"rango/tools/viper"

	"github.com/gin-gonic/gin"
)

func init() {
	viper.Init()
	logger.Init()
}

func Run() {
	engine := gin.Default()
	// 全局中间件注册
	engine.Use(middleware.CrossSetHeader, middleware.Init)
	// 注册路由
	ApiRouter(engine)
	UseRouter(engine)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	engine.Run("0.0.0.0:2020")
}
