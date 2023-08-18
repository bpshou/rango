package rds

import "context"

const LockPrefix = "distributed:locks:"

// 获取 Redis 分布式锁的key
func GetKey(key string) string {
	return LockPrefix + key
}

// 获取分布式锁
func Lock(key string) bool {
	ctx := context.Background()
	// 使用 SETNX 命令获取锁，如果返回值为 1 表示获取锁成功，返回 true 否则获取锁失败，返回 false
	ok, err := GetRedis().SetNX(ctx, GetKey(key), "1", 3600).Result()
	if err != nil {
		return false
	}
	return ok
}

// 释放分布式锁
func Unlock(key string) {
	// 使用 DEL 命令释放锁
	GetRedis().Del(context.Background(), GetKey(key)).Err()
}
