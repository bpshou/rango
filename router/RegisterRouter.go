package router

import (
	"rango/app/controller/api"
	"rango/app/controller/auth"
	"rango/app/controller/use"
	"rango/app/controller/user"
	"rango/middleware"

	docs "rango/docs"

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
		router.GET("/ping", api.Index{}.Ping)
		router.GET("/mysql", api.Index{}.Mysql)
		router.GET("/viper/config", api.Viper{}.ViperConfig)
		router.POST("/cmd", api.Command{}.Index)
		router.GET("/aes", api.Aes{}.Index)
		router.GET("/kafka", api.Kafka{}.Start)
		router.GET("/jwt", api.Jwt{}.Index)
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
	router = engine.Group("/use")
	{
		router.GET("/mongo", use.Mongo{}.Mongo)
		router.GET("/glog", use.Glog{}.Glog)
		router.GET("/logrus", use.Logrus{}.Logrus)
	}

	// 注册所有路由
	router = engine.Group("/route")
	{
		router.GET("/add", auth.Auth{}.AddGroupRoute(engine.Routes()))
	}
	// 文档
	docs.SwaggerInfo.BasePath = "/api"
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 静态资源
	engine.Static("/s", "./static")
}
