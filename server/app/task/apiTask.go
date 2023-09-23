package task

import (
	log "github.com/sirupsen/logrus"
)

type (
	ApiLogTask struct{}
)

// 轮训任务
func (j *ApiLogTask) Run() {
	for i := 0; i < 6; i++ {
		log.Debug("api task run", i)
	}
}
