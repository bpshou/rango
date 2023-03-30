package load

import (
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type TraceHook struct {
	TraceId string
}

func NewTraceIdHook(traceId string) logrus.Hook {
	hook := TraceHook{
		TraceId: traceId,
	}
	return &hook
}

func (hook *TraceHook) Fire(entry *logrus.Entry) error {
	entry.Data["trace"] = hook.TraceId
	return nil
}

func (hook *TraceHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func init() {
	// 纳秒 uuid
	trace := strconv.FormatInt(time.Now().UnixNano(), 10)
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)
	// 设置日志级别为debug
	logrus.SetLevel(logrus.DebugLevel)
	// 注册钩子
	logrus.AddHook(NewTraceIdHook(trace))
	// 成功
	logrus.Debug("init logrus Success !")
}
