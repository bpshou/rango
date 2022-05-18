package viper

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ImportConfig() {
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
			"omg": true,
			"err": err,
		}).Fatal("read config failed!")
	}

	log.WithFields(log.Fields{
		"status": true,
	}).Debug("Config init Success!")
}
