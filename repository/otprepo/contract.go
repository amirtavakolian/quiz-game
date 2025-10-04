package otprepo

import (
	"context"
	"time"
)

type OTPRepoContract interface {
	Set(ctx context.Context, key string, value int, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	TTL(ctx context.Context, key string) (time.Duration, error)
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, ttl time.Duration) (bool, error)
	Del(ctx context.Context, keys ...string) (int64, error)
}
