package rds

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func GetRedis() (redisClient *redis.Client) {
	// 实例化
	return redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.pass"),
		DB:       viper.GetInt("redis.db"),
	})
}
