package bucket

import (
	"context"
	"time"
)

// Repository interface contain methods to work with bucket storage
type Repository interface {
	Add(ctx context.Context, key string, capacity uint, rate time.Duration) error
	Reset(ctx context.Context, keys []string) error
	CleanStorage() error
}
