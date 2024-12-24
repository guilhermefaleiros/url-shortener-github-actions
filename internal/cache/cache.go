package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client redis.Client
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}) error {
	return c.client.Set(ctx, key, value, 0).Err()
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{client: *client}
}
