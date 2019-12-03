package antibruteforce

import "context"

// Usecase interface contain main antibruteforce methods
type Usecase interface {
	Check(ctx context.Context, login string, password string, ip string) error
	Reset(ctx context.Context, login string, ip string) error
	BlacklistAdd(ctx context.Context, subnet string) error
	BlacklistRemove(ctx context.Context, subnet string) error
	WhitelistAdd(ctx context.Context, subnet string) error
	WhitelistRemove(ctx context.Context, subnet string) error
}
