package task

import "github.com/robfig/cron/v3"

func Start() {
	// 精确到秒
	c := cron.New(
		cron.WithSeconds(), // 设置定时为秒
		cron.WithChain(
			cron.Recover(cron.DefaultLogger),             // panic 异常处理
			cron.DelayIfStillRunning(cron.DefaultLogger), // 一个任务执行一次
		),
	)

	c.AddJob("0 30 1 * * ?", &ApiLogTask{})  // 每30分的时候，每秒执行一次
	c.AddJob("00 30 22 * * ?", &KafkaTask{}) // 每天22:30执行一次

	// 启动
	c.Start()
	// 阻塞主线程停止
	select {}
}
