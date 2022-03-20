package database

import (
	config "github.com/berkayersoyy/go-products-example/pkg/utils/config"
	"github.com/go-redis/redis/v7"
)

type RedisClient interface {
	GetClient() *redis.Client
}

type redisClient struct {
	SingletonRedis *redis.Client
}

func ProvideRedisClient() RedisClient {
	return &redisClient{SingletonRedis: InitRedis()}
}
func (r *redisClient) GetClient() *redis.Client {
	return r.SingletonRedis
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
