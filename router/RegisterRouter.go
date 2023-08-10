package router

import (
	"rango/app/controller/api"
	"rango/app/controller/auth"
	"rango/app/controller/user"
	"rango/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRouter(engine *gin.Engine) {
	router := engine.Group("/api")
	{
		router.GET("/", api.Index{}.Index)
		router.GET("/totp", api.Totp{}.Secret)
		router.GET("/params/:id", api.Index{}.Params)
		router.GET("/mysql", api.Index{}.Mysql)
		router.GET("/mongo", api.Mongo{}.Mongo)
		router.GET("/kafka", api.Kafka{}.Start)
		router.GET("/secret/add", api.Secret{}.Create)
		router.GET("/secret/select", api.Secret{}.Decrypt)
		router.GET("/qrcode/create", api.Qrcode{}.Create)
	}
	router = engine.Group("/user")
	{
		router.POST("/reg", user.User{}.Register)
		router.Use(middleware.JWTAuth()).Use(middleware.RBAC()).GET("/login", user.User{}.Login)
		router.Use(middleware.JWTAuth()).Use(middleware.RBAC()).PUT("/edit", user.User{}.Edit)
		router.Use(middleware.JWTAuth()).Use(middleware.RBAC()).DELETE("/delete", user.User{}.Delete)
	}

	// 注册所有路由
	router = engine.Group("/route")
	{
		router.GET("/add", auth.Auth{}.AddGroupRoute(engine.Routes()))
	}
	// 文档
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 静态资源
	engine.Static("/s", "./static")
}
