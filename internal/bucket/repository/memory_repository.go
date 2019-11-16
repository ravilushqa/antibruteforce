package repository

import (
	"context"
	"gitlab.com/otus_golang/antibruteforce/internal/bucket/errors"
	"gitlab.com/otus_golang/antibruteforce/models"
	"time"
)

type MemoryBucketRepository struct {
	buckets map[string]*models.Bucket
}

func NewMemoryBucketRepository() *MemoryBucketRepository {
	return &MemoryBucketRepository{buckets: make(map[string]*models.Bucket, 1024)}
}

func (r *MemoryBucketRepository) Add(ctx context.Context, key string, capacity uint, rate time.Duration) error {
	b, ok := r.buckets[key]
	if !ok {
		b = &models.Bucket{
			Capacity:  capacity,
			Remaining: capacity - 1,
			Reset:     time.Now().Add(rate),
			Rate:      rate,
		}
		r.buckets[key] = b

		return nil
	}

	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	if time.Now().After(b.Reset) {
		b.Reset = time.Now().Add(b.Rate)
		b.Remaining = b.Capacity
	}

	if b.Remaining <= 0 {
		return errors.ErrBucketOverflow
	}

	b.Remaining -= 1

	return nil
}

func (r *MemoryBucketRepository) Reset(ctx context.Context, keys []string) error {
	for _, key := range keys {
		b, ok := r.buckets[key]
		if !ok {
			continue
		}

		b.Remaining = b.Capacity
	}
	return nil
}
