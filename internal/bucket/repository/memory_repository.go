package repository

import (
	"context"
	"github.com/ravilushqa/antibruteforce/internal/bucket/errors"
	"github.com/ravilushqa/antibruteforce/internal/bucket/models"
	"go.uber.org/zap"
	"sync"
	"time"
)

// clean duration in seconds
const cleanWaitDuration = 10

// MemoryBucketRepository is implementation of bucket repository interface
type MemoryBucketRepository struct {
	buckets map[string]*models.Bucket
	mutex   sync.Mutex
	l       *zap.Logger
}

// NewMemoryBucketRepository constructor for MemoryBucketRepository
func NewMemoryBucketRepository(logger *zap.Logger) *MemoryBucketRepository {
	m := &MemoryBucketRepository{buckets: make(map[string]*models.Bucket, 1024), l: logger}
	m.initCleaner()
	return m
}

func (r *MemoryBucketRepository) initCleaner() {
	go func() {
		for {
			time.Sleep(time.Duration(cleanWaitDuration) * time.Minute)
			err := r.CleanStorage()
			if err != nil {
				r.l.Error(err.Error())
			}
		}
	}()

}

// Add method is adding value to bucket, or creating it if its not created yet
func (r *MemoryBucketRepository) Add(ctx context.Context, key string, capacity uint, rate time.Duration) error {
	r.mutex.Lock()
	b, ok := r.buckets[key]
	if !ok {
		b = &models.Bucket{
			Capacity:  capacity,
			Remaining: capacity - 1,
			Reset:     time.Now().Add(rate),
			Rate:      rate,
		}
		r.buckets[key] = b

		r.mutex.Unlock()
		return nil
	}
	r.mutex.Unlock()

	if time.Now().After(b.Reset) {
		b.Reset = time.Now().Add(b.Rate)
		b.Remaining = b.Capacity
	}

	if b.Remaining == 0 {
		return errors.ErrBucketOverflow
	}

	b.Remaining--

	return nil
}

// Reset method resets buckets by keys
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

// CleanStorage us hard cleaning bucket storage
func (r *MemoryBucketRepository) CleanStorage() error {
	r.mutex.Lock()
	for k := range r.buckets {
		delete(r.buckets, k)
	}
	r.mutex.Unlock()

	return nil
}
