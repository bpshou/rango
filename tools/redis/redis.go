package redis

import (
	"github.com/go-redis/redis"
	// "fmt"
)

var rdb *redis.Client

func Instance() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.5.5:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// err := rdb.Set("key", "value", 0).Err()
	// fmt.Println(err)
	// if err != nil {
	// 	panic(err)
	// }
}
