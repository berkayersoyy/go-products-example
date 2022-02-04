package database

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
)

func InitRedis() *redis.Client {
	dsn := os.Getenv("REDIS_HOST")
	fmt.Println(dsn)
	if len(dsn) == 0 {
		dsn = "redis://redis:6379/"
	}
	client := redis.NewClient(&redis.Options{
		Addr: dsn,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
