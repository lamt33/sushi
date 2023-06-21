package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	DB *redis.Client
}

func NewRedisCache(o *redis.Options) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisCache{
		DB: rdb,
	}

}

func (pc RedisCache) Get(key string, dest interface{}) error {
	v, err := pc.DB.Get(key).Bytes()
	if err != nil {
		return &EntryNotFound{Key: key}
	}

	return json.Unmarshal(v, dest)
}

func (pc RedisCache) Set(key string, v interface{}, ttl time.Duration) error {
	p, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return pc.DB.Set(key, p, ttl).Err()
}
