package repo

import (
	"context"
	"time"
)

type Cache interface {
	// Put 放入
	Put(ctx context.Context, key string, value string, expire time.Duration) error
	// Get 获取
	Get(ctx context.Context, key string) (string, error)
}
