package errors

import "errors"

// ErrIPInBlackList error returning when requester in black list
var ErrIPInBlackList = errors.New("ip in black list")

// ErrWrongIP error returning IP is not valid
var ErrWrongIP = errors.New("is not a valid IP address")

// ErrLoginRequired error that said that login is required
var ErrLoginRequired = errors.New("login is required")

// ErrPasswordRequired error that said that password is required
var ErrPasswordRequired = errors.New("password is required")
