package container

import (
	"github.com/buemura/temp-storage/internal/application"
	"github.com/buemura/temp-storage/internal/infra/cache/redis"
	r "github.com/redis/go-redis/v9"
)

func LoadCacheStorage() *redis.RedisCacheStorage {
	cli := r.NewClient(&r.Options{
		Addr: "localhost:6379",
	})
	return redis.NewRedisCacheStorage(cli)
}

func LoadSessionService(cache *redis.RedisCacheStorage) *application.SessionService {
	ss := application.NewSessionService(cache)
	return ss
}
