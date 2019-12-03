package models

import (
	"sync"
	"time"
)

// Bucket model that contain attributes for implementing leaky bucket
type Bucket struct {
	Capacity  uint
	Remaining uint
	Reset     time.Time
	Rate      time.Duration
	Mutex     sync.Mutex
}
