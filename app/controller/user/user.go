package user

import (
	"rango/app/controller"
	"rango/app/service"

	"github.com/gin-gonic/gin"
)

type User struct {
	controller.Base
}

type (
	UserParams struct {
		Phone string `json:"phone"` // 手机号
		Email string `json:"email"` // 邮箱
		Code  int64  `json:"code"`  // 验证码
	}
)

// @Tags 		User模块
// @Summary 	用户登录/注册接口
// @Description 用户登录注册接口，如果以前没有注册过，则自动注册
// @Accept 		json
// @Produce 	json
// @Param     	data  body  UserParams  true  "请求入参"
// @Success 	200 {string} index
// @Router 		/user/login [get]
func (the User) Login(c *gin.Context) {
	token, err := service.ServiceGroupApp.UserService.Login()
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "token 获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

// @Tags 		User模块
// @Summary 	用户注册接口
// @Description 用户注册接口
// @Accept 		json
// @Produce 	json
// @Param     	data  body  UserParams  true  "请求入参"
// @Success 	200 {string} index
// @Router 		/user/register [post]
func (the User) Register(c *gin.Context) {
	token, err := service.ServiceGroupApp.UserService.Register()
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "token 获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

// @Tags 		User模块
// @Summary 	用户信息修改
// @Description 首页index入口，默认返回为一个index字符串
// @Accept 		json
// @Produce 	json
// @Success 	200 {string} index
// @Router 		/user/edit [put]
func (the User) Edit(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "修改成功",
	})
}

// @Tags 		User模块
// @Summary 	删除用户
// @Description 首页index入口，默认返回为一个index字符串
// @Accept 		json
// @Produce 	json
// @Success 	200 {string} index
// @Router 		/user/delete [delete]
func (the User) Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "删除成功",
	})
}
