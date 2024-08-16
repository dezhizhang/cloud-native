package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var RC *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	RC = &RedisCache{rdb: rdb}

}

// Put 设置数据
func (rc *RedisCache) Put(ctx context.Context, key string, value string, expire time.Duration) error {
	err := rc.rdb.Set(ctx, key, value, expire).Err()
	return err
}

// Get 获取数据
func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.rdb.Get(ctx, key).Result()
	return result, err
}
