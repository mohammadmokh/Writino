package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"gitlab.com/gocastsian/writino/config"
)

type RedisStore struct {
	client *redis.Client
}

func New(ctx context.Context, cfg config.RedisCfg) (RedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	res := client.Ping(ctx)

	return RedisStore{
		client: client,
	}, res.Err()
}
