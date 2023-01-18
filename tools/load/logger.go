package load

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)
	// 设置日志级别为debug
	logrus.SetLevel(logrus.DebugLevel)
	// 成功
	logrus.Debug("init logrus Success !")
}
