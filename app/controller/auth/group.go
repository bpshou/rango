package auth

import (
	"rango/app/controller"
	"rango/tools"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	controller.Base
}

// @Tags 		Auth模块
// @Summary 	注册所有路由接口
// @Description 调用后，自动将路由全部注册到表中，添加的super组，拥有所有权限
// @Accept 		json
// @Produce 	json
// @Param     	data  body  gin.RoutesInfo  true  "请求入参"
// @Success 	200 {string} index
// @Router 		/route/add [get]
func (the Auth) AddGroupRoute(routes gin.RoutesInfo) func(c *gin.Context) {
	return func(c *gin.Context) {
		for _, v := range routes {
			enforcer, err := tools.GetEnforcer()
			if err != nil {
				return
			}

			// 自动补全
			hasPolicy := enforcer.HasPolicy("super", v.Path, v.Method)
			if !hasPolicy {
				enforcer.AddGroupingPolicy("super", v.Path, v.Method)
			}
		}

		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "所有路由均已注册到指定组",
		})
	}
}
