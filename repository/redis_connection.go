package repository

import (
	"context"
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/configloader"
	"github.com/redis/go-redis/v9"
)

func NewRedisConnection(ctx context.Context) *redis.Client {

	cfgLoader := configloader.NewConfigLoader()
	k := cfgLoader.SetPrefix("APP_").SetDivider("_").SetDelimiter(".").Build()

	rdb := redis.NewClient(&redis.Options{
		Addr:     k.String("redis.host") + ":" + k.String("redis.port"),
		Password: k.String("redis.password"),
		DB:       k.Int("redis.db"),
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err.Error())
	}

	return rdb
}
