package middleware

import (
	"rango/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("authorization")

		logrus.Debug(token)
		if token == "" {
			c.JSON(400, gin.H{"error": "未登录或非法访问"})
			c.Abort()
			return
		}

		jwtTools := tools.NewJWT()
		claims, err := jwtTools.ParseToken(token)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims.Uid <= 0 {
			c.JSON(400, gin.H{"error": "user info error"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
