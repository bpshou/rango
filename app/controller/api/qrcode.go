package api

import (
	"rango/app/controller"
	"rango/app/service"

	"github.com/gin-gonic/gin"
)

type Qrcode struct {
	controller.Base
}

func (the Qrcode) Create(c *gin.Context) {
	// 内容
	content := c.Query("content")

	qr, err := service.CreateQrcode(content)
	if err != nil {
		c.String(400, "失败")
		return
	}
	c.String(200, "成功\n%s", qr)
}
