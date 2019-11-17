package errors

import "errors"

var ErrIpInBlackList = errors.New("ip in black list")
var ErrWrongIp = errors.New("is not a valid IP address")
var ErrLoginRequired = errors.New("login is required")
var ErrPasswordRequired = errors.New("password is required")
