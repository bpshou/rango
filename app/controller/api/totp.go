package api

import (
	"rango/app/controller"
	"rango/app/service"

	"github.com/gin-gonic/gin"
)

type Totp struct {
	controller.Base
}

func (this Totp) Secret(c *gin.Context) {
	// aes
	google_secret := "/N0IMuSTHTPWoy8H096B8ipENjP0oASLBmnBFapJmGo="
	// aes
	vpn_secret := "cBg/F/TFpuVlzmcWoRzH/Kn/78tTQRiYyAkK8Un6y7cv2fnZ0n75VuTSm+loBUNS"
	// aes
	arm_secret := "VBfkXvv53nN3CROGgtxG97jt/5NPxRSO+XkjD3bfFq0="
	google, _ := service.CreateTotpCode(google_secret, 30)
	vpn, _ := service.CreateTotpCode(vpn_secret, 60)
	arm, _ := service.CreateTotpCode(arm_secret, 30)
	c.String(200, "谷歌认证 OTP is: %s", google)
	c.String(200, "\n")
	c.String(200, google)
	c.String(200, "\n")
	c.String(200, "远程VPN OTP is: %s", vpn)
	c.String(200, "\n")
	c.String(200, vpn)
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
