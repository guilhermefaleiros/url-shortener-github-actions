package storage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

func NewRedisClient(ctx context.Context) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
}
