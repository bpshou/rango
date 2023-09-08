package main

import (
	"rango/router"
)

// @title           rango 服务接口 Swagger 文档
// @version         1.0
// @description     rango服务端接口文档，注意模块名称
// @host      		localhost:2020

// @securityDefinitions.basic  BasicAuth
func main() {
	router.Run()
}
