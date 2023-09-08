package claims

import (
	"rango/tools"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) *tools.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		return nil
	} else {
		waitUse := claims.(*tools.CustomClaims)
		return waitUse
	}
}

func GetUserId(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); !exists {
		return 0
	} else {
		waitUse := claims.(*tools.CustomClaims)
		return waitUse.Uid
	}
}

func GetPhone(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		return ""
	} else {
		waitUse := claims.(*tools.CustomClaims)
		return waitUse.Phone
	}
}

func GetNickname(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		return ""
	} else {
		waitUse := claims.(*tools.CustomClaims)
		return waitUse.NickName
	}
}
