// pkg/config/cache.go

package config

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetRedisClient() *redis.Client {
	return RedisClient
}

func SetRedis(ctx context.Context, key string, value interface{}) error {
	return RedisClient.Set(ctx, key, value, 0).Err()
}

func GetRedis(ctx context.Context, key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

func DelRedis(ctx context.Context, key string) error {
	return RedisClient.Del(ctx, key).Err()
}
