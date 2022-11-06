package Cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

type RedisCache struct {
	redis *redis.Client
	TTL   time.Duration
}

func NewRedisCache(addr string, password string, db int, ttl int) *RedisCache {
	return &RedisCache{
		redis: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
		TTL: time.Second * time.Duration(ttl),
	}
}

func (r RedisCache) Store(s string, s2 []byte) error {
	return r.redis.Set(context.Background(), s, s2, r.TTL).Err()
}

func (r RedisCache) Get(s string) ([]byte, error) {
	get := r.redis.Get(context.Background(), s)
	return get.Bytes()
}

func (r RedisCache) Clear(s string) error {
	return r.redis.Del(context.Background(), s).Err()
}
