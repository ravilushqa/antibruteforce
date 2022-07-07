package models

import (
	"sync"
)

// Bucket model that contain attributes for implementing leaky bucket
type Bucket struct {
	Capacity  uint
	Remaining int32
	Mutex     sync.Mutex
}
