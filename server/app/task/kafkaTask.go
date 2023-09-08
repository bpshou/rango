package task

import (
	"fmt"
	"rango/tools/kafka"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func KafkaTask() {
	var wg sync.WaitGroup
	consumer := kafka.Consumer()

	// topic
	topic := viper.GetString("kafka.topic.task")
	// 通过topic获取到所有的分区
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("Failed to get the list of partition: ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList { // 遍历所有的分区
		partConsumer, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest) // 针对每个分区创建一个分区消费者
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) { // 为每个分区开一个go协程取值
			for msg := range partConsumer.Messages() { // 阻塞直到有值发送过来，然后再继续等待
				fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			defer partConsumer.AsyncClose()
			wg.Done()
		}(partConsumer)
	}
	wg.Wait()
	consumer.Close()
}
