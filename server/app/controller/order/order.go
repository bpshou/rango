package order

import (
	"rango/app/controller"
	"rango/app/response"

	"github.com/gin-gonic/gin"
)

type Order struct {
	controller.Base
}

type (
	OrderParams struct {
		Phone string `json:"phone"`           // 手机号
		Email string `json:"email,omitempty"` // 邮箱
		Code  string `json:"code"`            // 验证码
	}
)

// @Tags 		Order模块
// @Summary 	生成订单
// @Description 用户购买账号，生成订单供用户支付
// @Accept 		json
// @Produce 	json
// @Success 	200 {string} index
// @Router 		/user/edit [put]
func (the Order) CreateOrder(c *gin.Context) {
	response.OkWithMessage("创建成功", c)
}

// @Tags 		Order模块
// @Summary 	支付订单回调
// @Description 订单付款以后，支付回调接口
// @Accept 		json
// @Produce 	json
// @Success 	200 {string} index
// @Router 		/user/edit [put]
func (the Order) OrderCallback(c *gin.Context) {
	response.OkWithMessage("回调成功", c)
}
