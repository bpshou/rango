package api

import (
	"encoding/base64"
	"rango/app/controller"
	"rango/app/service"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

type Totp struct {
	controller.Base
}

type TotpParams struct {
	Secret string `form:"secret" json:"secret" xml:"secret"`
}

func (the Totp) Secret(c *gin.Context) {
	// 解析参数
	var params TotpParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(200, gin.H{
			"code":    4001,
			"message": err.Error(),
		})
		return
	}

	// 数据
	secret := params.Secret
	if secret != "" {
		secretByte, _ := base64.StdEncoding.DecodeString(secret)
		secret = string(secretByte)
	}

	// aes
	googleByte, _ := service.Decrypt(viper.GetString("google.secret"), secret)
	// aes
	vpnByte, _ := service.Decrypt(viper.GetString("vpn.secret"), secret)
	// aes
	armByte, _ := service.Decrypt(viper.GetString("arm.secret"), secret)

	google, _ := service.CreateTotpCode(string(googleByte), 30)
	vpn, _ := service.CreateTotpCode(string(vpnByte), 60)
	arm, _ := service.CreateTotpCode(string(armByte), 30)

	// 输出
	c.String(200, "谷歌认证 OTP is: %s", google)
	c.String(200, "\n")
	c.String(200, google)
	c.String(200, "\n")
	c.String(200, "远程VPN OTP is: %s", vpn)
	c.String(200, "\n")
	c.String(200, vpn)
	c.String(200, "\n")
	c.String(200, "远程ARM OTP is: %s", arm)
	c.String(200, "\n")
	c.String(200, arm)
	// c.JSON(200, gin.H{
	// 	"code":    200,
	// 	"message": "secret",
	// 	"google":  google,
	// 	"vpn":  vpn,
	// })
}
