package utils

import (
	"os"
	"github.com/sirupsen/logrus"
)

func LogInit() {

	logrus.SetOutput(os.Stdout)
	// 设置logrus实例输出格式为json
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Debug("test debug logrus")
}
