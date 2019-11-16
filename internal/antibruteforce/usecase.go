package antibruteforce

import "context"

type Usecase interface {
	Check(ctx context.Context, login string, password string, ip string) (ok bool, err error)
}
