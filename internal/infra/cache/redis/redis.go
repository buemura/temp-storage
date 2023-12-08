package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCacheStorage struct {
	r *redis.Client
}

func NewRedisCacheStorage(r *redis.Client) *RedisCacheStorage {
	return &RedisCacheStorage{
		r: r,
	}
}

func (c *RedisCacheStorage) Connect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func (c *RedisCacheStorage) Close() error {
	return c.r.Close()
}

func (c *RedisCacheStorage) Get(key string) string {
	ctx := context.Background()
	val, _ := c.r.Get(ctx, key).Result()
	return val
}

func (c *RedisCacheStorage) Set(key string, data string, ttl int) error {
	ctx := context.Background()
	return c.r.Set(ctx, key, data, time.Duration(ttl)*time.Minute).Err()
}
