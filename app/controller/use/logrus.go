package use

import (
	"rango/app/controller"

	"github.com/gin-gonic/gin"

	"os"

	log "github.com/sirupsen/logrus"
)

type Logrus struct {
	controller.Base
}

func (the Logrus) Logrus(c *gin.Context) {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.DebugLevel)

	log.WithFields(log.Fields{
		"name": "zhangsan",
		"age":  18,
	}).Debug("A debug message")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"database": "mysql",
		"port":     3306,
	}).Error("Mysql connect error!")

	// 使用 Fatal 会中断程序
	// log.WithFields(log.Fields{
	// 	"omg":    true,
	// 	"number": 100,
	// }).Fatal("The ice breaks!")

	c.JSON(200, gin.H{
		"code":    200,
		"message": "Use logrus success !",
	})
}
