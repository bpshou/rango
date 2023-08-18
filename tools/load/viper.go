package load

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 加载viper，注意路径
func LoadViper(path string) {
	// 设置读取配置文件类型
	viper.SetConfigType("yml")
	// 设置viper配置读取目录
	viper.AddConfigPath(path + "config")
	viper.AddConfigPath(path + "config/dev")
	viper.AddConfigPath(path + "config/test")
	viper.AddConfigPath(path + "config/online")

	// 读取配置
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"method": "viper.ReadInConfig",
			"err":    err,
		}).Warn("Read config failed!")
	}

	// 读取redis配置
	// viper.SetConfigName("redis")
	// viper.SetConfigName("kafka")
	viper.SetConfigName("secret")
	viper.SetConfigName("sdk")
	viper.MergeInConfig()

	log.Debug("init viper Success !")
}
