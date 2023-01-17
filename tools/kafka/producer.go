package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// kafka客户端
func Client() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 写到随机分区中，我们默认设置32个分区
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	broker := viper.GetString("kafka.broker")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		logrus.Error("Producer closed error", err)
	}

	return client
}

// 通过kafka客户端发送消息
func SendMsg(client sarama.SyncProducer, message *sarama.ProducerMessage) {
	// 发送消息
	partition, offset, err := client.SendMessage(message)
	if err != nil {
		logrus.Warn("Send message failed", err)
		return
	}
	defer client.Close()

	logrus.WithFields(logrus.Fields{
		"partition": partition, // 分区
		"offset":    offset,    // 索引
	}).Debug("Send message Success")
}

// 通过kafka客户端发送消息
func SendTopicMsg(client sarama.SyncProducer, topic string, message string) {
	// 消息体
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(message)

	// 发送消息
	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		logrus.Warn("Send message failed", err)
		return
	}
	defer client.Close()

	logrus.WithFields(logrus.Fields{
		"partition": partition, // 分区
		"offset":    offset,    // 索引
	}).Debug("Send message Success")
}
