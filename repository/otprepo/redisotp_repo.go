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

func (redis RedisOTPRepo) Get(ctx context.Context, key string) (string, error) {
	val, err := redis.Client.Get(ctx, key).Result()
	return val, err
}

func (redis RedisOTPRepo) TTL(ctx context.Context, key string) (time.Duration, error) {
	val, err := redis.Client.TTL(ctx, key).Result()
	return val, err
}

func (redis RedisOTPRepo) Incr(ctx context.Context, key string) (int64, error) {
	val, err := redis.Client.Incr(ctx, key).Result()
	return val, err
}

func (redis RedisOTPRepo) Expire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	val, err := redis.Client.Expire(ctx, key, ttl).Result()
	return val, err
}

func (redis RedisOTPRepo) Del(ctx context.Context, keys ...string) (int64, error) {
	deleted, err := redis.Client.Del(ctx, keys...).Result()
	return deleted, err
}
