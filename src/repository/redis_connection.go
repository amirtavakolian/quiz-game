package repository

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func NewRedisConnection(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: func() string {
			return fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
		}(),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: func() int {
			num, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
			return num
		}(),
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err.Error())
	}

	return rdb
}
