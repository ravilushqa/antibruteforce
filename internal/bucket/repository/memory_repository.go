package repository

import (
	"context"
	"sync/atomic"
	"time"

	"go.uber.org/zap"

	"github.com/ravilushqa/antibruteforce/internal/bucket/errors"
	"github.com/ravilushqa/antibruteforce/internal/bucket/models"
)

// MemoryBucketRepository is implementation of bucket repository interface
type MemoryBucketRepository struct {
	buckets map[string]*models.Bucket
	l       *zap.Logger
}

// NewMemoryBucketRepository constructor for MemoryBucketRepository
func NewMemoryBucketRepository(logger *zap.Logger) *MemoryBucketRepository {
	m := &MemoryBucketRepository{buckets: make(map[string]*models.Bucket, 1024), l: logger}
	return m
}

// Add method is adding value to bucket, or creating it if its not created yet
func (r *MemoryBucketRepository) Add(ctx context.Context, key string, capacity uint, rate time.Duration) error {
	b, ok := r.buckets[key]
	if !ok {
		b = &models.Bucket{Capacity: capacity}
		atomic.AddInt32(&b.Remaining, int32(capacity-1))
		go func() {
			time.Sleep(rate)
			atomic.AddInt32(&b.Remaining, 1)
		}()
		r.buckets[key] = b

		return nil
	}

	if atomic.LoadInt32(&b.Remaining) == 0 {
		return errors.ErrBucketOverflow
	}

	atomic.AddInt32(&b.Remaining, -1)
	go func() {
		time.Sleep(rate)
		atomic.AddInt32(&b.Remaining, 1)
	}()
	return nil
}

// Reset method resets buckets by keys
func (r *MemoryBucketRepository) Reset(ctx context.Context, keys []string) error {
	for _, key := range keys {
		b, ok := r.buckets[key]
		if !ok {
			continue
		}

		atomic.StoreInt32(&b.Remaining, int32(b.Capacity))
	}
	return nil
}
