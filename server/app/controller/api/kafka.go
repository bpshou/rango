package api

import (
	"rango/app/controller"
	"rango/tools/kafka"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Kafka struct {
	controller.Base
}

func (the Kafka) Start(c *gin.Context) {
	// 使用客户端
	client := kafka.Client()
	// topic
	topic := viper.GetString("kafka.topic.task")
	// 发送消息
	kafka.SendTopicMsg(client, topic, "this is message")
}
