package middleware

import (
	"fmt"
	"rango/common/claims"
	"rango/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	_ "github.com/go-sql-driver/mysql"
)

func RBAC() gin.HandlerFunc {
	return func(c *gin.Context) {
		enforcer, err := tools.GetEnforcer()
		if err != nil {
			c.Abort()
			return
		}

		userId := claims.GetUserId(c)
		// 请求路径
		path := c.Request.URL.Path
		// 请求方法
		act := c.Request.Method

		// 校验
		ok, err := enforcer.Enforce(cast.ToString(userId), path, act)
		if err != nil {
			logrus.Fatalf("error, detail: %s", err)
			c.Abort()
			return
		}

		fmt.Println(cast.ToString(userId), path, act)
		fmt.Println(ok)

		if !ok {
			c.Abort()
			return
		}

		c.Next()
	}
}
