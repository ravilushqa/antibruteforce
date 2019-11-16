package repository

import (
	"context"
	"testing"
	"time"
)

var memoryBucketRepository = NewMemoryBucketRepository()
var rate = 100 * time.Millisecond

func TestMemoryBucketRepository_Add(t *testing.T) {
	type args struct {
		ctx      context.Context
		key      string
		capacity uint
		rate     time.Duration
	}
	tests := []struct {
		sleep   time.Duration
		name    string
		args    args
		wantErr bool
	}{
		{
			sleep: 0,
			name:  "testing leaky bucket, first iteration, capacity 2",
			args: args{
				ctx:      context.Background(),
				key:      "test",
				capacity: 2,
				rate:     rate,
			},
			wantErr: false,
		},
		{
			sleep: 0,
			name:  "testing leaky bucket, second iteration, capacity 2",
			args: args{
				ctx:      context.Background(),
				key:      "test",
				capacity: 2,
				rate:     0,
			},
			wantErr: false,
		},
		{
			sleep: 0,
			name:  "testing leaky bucket, another key",
			args: args{
				ctx:      context.Background(),
				key:      "test-another-key",
				capacity: 2,
				rate:     0,
			},
			wantErr: false,
		},
		{
			sleep: 0,
			name:  "testing leaky bucket, third iteration, capacity 2",
			args: args{
				ctx:      context.Background(),
				key:      "test",
				capacity: 2,
				rate:     0,
			},
			wantErr: true,
		},
		{
			sleep: rate,
			name:  "testing leaky bucket, third iteration, capacity 2",
			args: args{
				ctx:      context.Background(),
				key:      "test",
				capacity: 2,
				rate:     0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		time.Sleep(tt.sleep)

		t.Run(tt.name, func(t *testing.T) {

			if err := memoryBucketRepository.Add(tt.args.ctx, tt.args.key, tt.args.capacity, tt.args.rate); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
