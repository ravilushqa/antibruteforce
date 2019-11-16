package bucket

import (
	"context"
	"time"
)

type Repository interface {
	Add(ctx context.Context, key string, capacity uint, rate time.Duration) error
}
