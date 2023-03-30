package middleware

import (
	"rango/tools/load"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 初始化配置
func Init(c *gin.Context) {
	// 校验权限

	c.Next()
}

// 生成 traceid
func AddTrace(c *gin.Context) {
	// 纳秒 uuid
	trace := strconv.FormatInt(time.Now().UnixNano(), 10)
	// 注册钩子
	logrus.AddHook(load.NewTraceIdHook(trace))
}
