package task

import "github.com/robfig/cron/v3"

func Start() {
	// 精确到秒
	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/1 30 * * * ?", ApiLogTask) // 每30分的时候，每秒执行一次
	c.AddFunc("00 30 22 * * ?", KafkaTask)  // 每天22:30执行一次

	// 启动
	c.Start()
	// 阻塞主线程停止
	select {}
}
