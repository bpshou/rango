package service

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "fmt"
)

func ViperConfig(c *gin.Context)  {
    host := ""
    if viper.IsSet("host") {
        host = viper.Get("host").(string)
    }
    fmt.Println("redis host after: ", host)
}