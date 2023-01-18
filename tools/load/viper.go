package load

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 加载viper，注意路径
func LoadViper(path string) {
	// 设置读取配置文件类型
	viper.SetConfigType("json")
	// 设置viper配置读取目录
	viper.AddConfigPath(path + "app/config")
	viper.AddConfigPath(path + "app/config/dev")
	viper.AddConfigPath(path + "app/config/test")
	viper.AddConfigPath(path + "app/config/prod")

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
	viper.SetConfigName("kafka")
	viper.SetConfigName("common")
	viper.MergeInConfig()

	log.Debug("init viper Success !")
}
