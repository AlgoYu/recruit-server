package redis

import "github.com/go-redis/redis"

var Rdb *redis.Client

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
}
