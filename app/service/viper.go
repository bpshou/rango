package service

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "fmt"
)

func ViperConfig(c *gin.Context)  {
    host := viper.Get("host")
    fmt.Println("redis host after: ", host)
}