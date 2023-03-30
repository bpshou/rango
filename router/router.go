package router

import (
	"rango/app/task"
	"rango/middleware"

	// 隐式导包实现初始化
	"rango/tools/load" // 存在隐式加载，注意有路径问题

	"github.com/gin-gonic/gin"
)

func Run() {
	// 加载配置
	load.LoadViper("./")
	// 启动任务
	task.Start()

	engine := gin.Default()
	// 全局中间件注册
	engine.Use(middleware.AddTrace, middleware.CrossSetHeader, middleware.Init)
	// 注册路由
	ApiRouter(engine)
	UseRouter(engine)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	engine.Run("0.0.0.0:2020")
}
