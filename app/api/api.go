package api

import (
	"rango/utils/logger"
	"rango/router"
)

func Run() {
	logger.Init()
	router.Run()
}
