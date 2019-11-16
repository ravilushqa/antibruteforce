package models

import (
	"sync"
	"time"
)

type Bucket struct {
	Capacity  uint
	Remaining uint
	Reset     time.Time
	Rate      time.Duration
	Mutex     sync.Mutex
}
