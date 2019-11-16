package antibruteforce

import "context"

type Repository interface {
	BlacklistAdd(ctx context.Context, subnet string) error
	BlacklistRemove(ctx context.Context, subnet string) error
	WhitelistAdd(ctx context.Context, subnet string) error
	WhitelistRemove(ctx context.Context, subnet string) error
}
