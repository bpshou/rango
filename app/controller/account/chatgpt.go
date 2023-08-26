package account

import (
	"fmt"
	"rango/app/controller"
	"rango/app/response"
	"rango/app/service"
	"rango/common/claims"

	"github.com/gin-gonic/gin"
)

type Account struct {
	controller.Base
}

type (
	AccountParams struct {
		Phone string `json:"phone"`           // 手机号
		Email string `json:"email,omitempty"` // 邮箱
		Code  string `json:"code"`            // 验证码
	}
)

// @Tags 		Account模块
// @Summary 	下发账号
// @Description 用户购买成功后，下发chatGPT账号
// @Accept 		json
// @Produce 	json
// @Success 	200 {string} index
// @Router 		/user/edit [post]
func (the Account) PutAccount(c *gin.Context) {
	// 用户ID
	userId := claims.GetUserId(c)

	fmt.Println(userId)

	token, err := service.ServiceGroupApp.AccountService.PutAccount(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(map[string]any{
		"token": token,
	}, "下发成功", c)
}

// @Tags 		Account模块
// @Summary 	chatgpt入口
// @Description 登录成功以后，下发的chatgpt入口地址，用户可以直接使用
// @Accept 		json
// @Produce 	json
// @Success 	200 {string} index
// @Router 		/user/edit [post]
func (the Account) PutEndpoint(c *gin.Context) {
	response.OkWithDetailed(map[string]any{
		"url": "http://www.baidu.com",
	}, "下发成功", c)
}
