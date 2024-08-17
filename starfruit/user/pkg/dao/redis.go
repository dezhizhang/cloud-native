package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"starfruit.top/user/config"
	"time"
)

var RC *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(config.AppConfig.ReadRedisConfig())
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
