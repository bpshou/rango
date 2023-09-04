package user

import (
	"rango/app/controller"
	"rango/app/response"
	"rango/app/service"

	"github.com/gin-gonic/gin"
)

type User struct {
	controller.Base
}

type (
	UserParams struct {
		Phone string `json:"phone"`           // 手机号
		Email string `json:"email,omitempty"` // 邮箱
		Code  string `json:"code"`            // 验证码
	}

	SmsParams struct {
		Phone string `json:"phone"` // 手机号
	}
)

// @Tags 		User模块
// @Summary 	用户登录/注册接口
// @Description 用户登录注册接口，如果以前没有注册过，则自动注册
// @Accept 		json
// @Produce 	json
// @Param     	data  body  UserParams  true  "请求入参"
// @Success 	200 {object} response.Response{msg=string, data=map[string]string}
// @Router 		/user/login [post]
func (the User) Login(c *gin.Context) {
	// 获取json参数
	var params UserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailWithMessage("参数错误："+err.Error(), c)
		return
	}

	if params.Phone == "" || params.Code == "" {
		response.FailWithMessage("参数为空", c)
		return
	}

	if !service.Group.SmsService.CheckCode(params.Phone, params.Code) {
		response.FailWithMessage("短信验证码错误", c)
		return
	}

	token, err := service.Group.UserService.Login(params.Phone)
	if err != nil {
		response.FailWithMessage("登录失败", c)
		return
	}
	response.OkWithDetailed(gin.H{
		"token": token,
	}, "登录成功", c)
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
	// 获取json参数
	var params UserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailWithMessage("参数错误："+err.Error(), c)
		return
	}

	if !service.Group.SmsService.CheckCode(params.Phone, params.Code) {
		response.FailWithMessage("短信验证码错误", c)
		return
	}

	token, err := service.Group.UserService.Login(params.Phone)
	if err != nil {
		response.FailWithMessage("注册失败", c)
		return
	}
	response.OkWithDetailed(gin.H{
		"token": token,
	}, "注册成功", c)
}

// @Tags 		User模块
// @Summary 	用户信息修改
// @Description 首页index入口，默认返回为一个index字符串
// @Accept 		json
// @Produce 	json
// @Success 	200 {string} index
// @Router 		/user/edit [put]
func (the User) Edit(c *gin.Context) {
	response.OkWithMessage("修改成功", c)
}

// @Tags 		User模块
// @Summary 	删除用户
// @Description 首页index入口，默认返回为一个index字符串
// @Accept 		json
// @Produce 	json
// @Success 	200 {string} index
// @Router 		/user/delete [delete]
func (the User) Delete(c *gin.Context) {
	response.OkWithMessage("删除成功", c)
}

// @Tags 		User模块
// @Summary 	短信发送
// @Description 用户登录发送短信的接口
// @Accept 		json
// @Produce 	json
// @Param     	data  body  SmsParams  true  "请求入参"
// @Success 	200 {string} index
// @Router 		/user/sms/send [get]
func (the User) SmsSend(c *gin.Context) {
	// 获取json参数
	var params SmsParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailWithMessage("参数错误："+err.Error(), c)
		return
	}

	_, err := service.Group.SmsService.Send(params.Phone)
	if err != nil {
		response.FailWithMessage("短信发送失败", c)
		return
	}
	response.OkWithMessage("短信发送成功", c)
}
