package database

import (
	config "github.com/berkayersoyy/go-products-example/pkg/utils/config"
	"github.com/go-redis/redis/v7"
)

var singletonRedis *redis.Client

func GetRedisClient() *redis.Client {
	if singletonRedis == nil {
		singletonRedis = InitRedis()
	}
	return singletonRedis
}

func InitRedis() *redis.Client {
	conf, err := config.LoadConfig("./")
	if err != nil {
		panic(err)
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
