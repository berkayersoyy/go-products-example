package database

import (
	config "github.com/berkayersoyy/go-products-example/pkg/utils/config"
	"github.com/go-redis/redis/v7"
)

type redisClient struct {
	SingletonRedis *redis.Client
}

var redisclient redisClient

func GetRedisClient() redisClient {
	if redisclient.SingletonRedis == nil {
		redisclient.SingletonRedis = InitRedis()
	}
	return redisclient
}

func InitRedis() *redis.Client {
	conf, err := config.LoadConfig("./")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: conf.RedisHost,
	})
	_, err = client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
