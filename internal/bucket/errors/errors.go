package errors

import (
	"errors"
)

// ErrBucketOverflow error returning when bucket is overflow
var ErrBucketOverflow = errors.New("bucket is overflow")
