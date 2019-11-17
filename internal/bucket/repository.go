package bucket

import (
	"context"
	"time"
)

type Repository interface {
	Add(ctx context.Context, key string, capacity uint, rate time.Duration) error
	Reset(ctx context.Context, keys []string) error
	CleanStorage() error
}
