package viper

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("db")
	viper.SetConfigName("redis")
	// viper.SetConfigType("json")
	viper.AddConfigPath("./app/config")
	viper.AddConfigPath("./app/config/dev")
	viper.AddConfigPath("./app/config/test")
	viper.AddConfigPath("./app/config/prod")

	err := viper.ReadInConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"method": "viper.ReadInConfig",
			"err":    err,
		}).Fatal("Read config failed!")
	}
}

func Init() {
	log.Debug("Config viper init Success !")
}
