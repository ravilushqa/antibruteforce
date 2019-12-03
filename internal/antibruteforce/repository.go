package antibruteforce

import "context"

// Repository interface contain methods to work with storage
type Repository interface {
	BlacklistAdd(ctx context.Context, subnet string) error
	BlacklistRemove(ctx context.Context, subnet string) error
	WhitelistAdd(ctx context.Context, subnet string) error
	WhitelistRemove(ctx context.Context, subnet string) error
	FindIPInList(ctx context.Context, ip string) (string, error)
}
