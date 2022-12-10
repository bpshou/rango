package viper

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// 设置读取配置文件类型
	viper.SetConfigType("json")
	// 设置viper配置读取目录
	viper.AddConfigPath("./app/config")
	viper.AddConfigPath("./app/config/dev")
	viper.AddConfigPath("./app/config/test")
	viper.AddConfigPath("./app/config/prod")

	// 读取数据库配置
	viper.SetConfigName("db")

	err := viper.ReadInConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"method": "viper.ReadInConfig",
			"err":    err,
		}).Warn("Read config failed!")
	}

	// 读取redis配置
	viper.SetConfigName("redis")
	viper.MergeInConfig()

	log.Debug("init viper Success !")
}
