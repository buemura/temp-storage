package cache

type CacheStorage interface {
	Connect() *CacheStorage
	Close() error
	Get(key string) string
	Set(key string, data string, ttl int) error
}
