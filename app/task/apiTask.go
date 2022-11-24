package task

import (
	log "github.com/sirupsen/logrus"
)

func ApiLogTask() {
	for i := 0; i < 6; i++ {
		log.Debug("api task run", i)
	}
}
