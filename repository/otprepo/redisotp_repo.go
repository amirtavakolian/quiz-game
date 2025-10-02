package otprepo

import (
	"context"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/repository"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisOTPRepo struct {
	Client *redis.Client
}

func NewRedisOTPRepo() RedisOTPRepo {
	ctx := context.Background()
	return RedisOTPRepo{Client: repository.NewRedisConnection(ctx)}
}

func (redis RedisOTPRepo) Set(ctx context.Context, key string, value int, ttl time.Duration) error {
	if err := redis.Client.Set(ctx, key, value, ttl).Err(); err != nil {
		return err
	}
	return nil
}
