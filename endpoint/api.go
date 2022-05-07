package endpoint

import (
    "rango/app/config"
    "rango/utils/logger"
    "rango/router"
)

func Run() {
    config.Init()
    logger.Init()
    router.Run()
}
