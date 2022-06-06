package redis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func Instance() *redis.Client {
	var rdb *redis.Client
	// 实例化
	rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.pass"), // no password set
		DB:       viper.GetInt("redis.db"),      // use default DB
	})
	return rdb
}
