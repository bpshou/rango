package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// kafka消费者
func Consumer() sarama.Consumer {
	broker := viper.GetString("kafka.broker")
	// 消费者
	consumer, err := sarama.NewConsumer([]string{broker}, nil)
	if err != nil {
		logrus.Error("Consumer start error", err)
	}

	return consumer
}
