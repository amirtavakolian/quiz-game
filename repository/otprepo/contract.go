package otprepo

import (
	"context"
	"time"
)

type OTPContract interface {
	Set(ctx context.Context, key string, value int, ttl time.Duration) error
}
