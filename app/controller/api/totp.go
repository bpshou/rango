package api

import (
	"rango/app/controller"
	"rango/app/service"

	"github.com/gin-gonic/gin"
)

type Totp struct {
	controller.Base
}

func (this Totp) Secret(c *gin.Context)  {
	// aes
	google_secret := "/N0IMuSTHTPWoy8H096B8ipENjP0oASLBmnBFapJmGo="
	// aes
	vpn_secret := "cBg/F/TFpuVlzmcWoRzH/Kn/78tTQRiYyAkK8Un6y7cv2fnZ0n75VuTSm+loBUNS"
	google, _ := service.CreateTotpCode(google_secret, 30)
	vpn, _ := service.CreateTotpCode(vpn_secret, 60)
	c.String(200, "谷歌认证 OTP is: %s", google)
	c.String(200, "\n")
	c.String(200, google)
	c.String(200, "\n")
	c.String(200, "远程VPN OTP is: %s", vpn)
	c.String(200, "\n")
	c.String(200, vpn)
	// c.JSON(200, gin.H{
	// 	"code":    200,
	// 	"message": "secret",
	// 	"google":  google,
	// 	"vpn":  vpn,
	// })
}