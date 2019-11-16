package repository

import "context"

type psqlAntibruteforceRepository struct {
}

func NewPsqlAntibruteforceRepository() *psqlAntibruteforceRepository {
	return &psqlAntibruteforceRepository{}
}

func (p psqlAntibruteforceRepository) BlacklistAdd(ctx context.Context, subnet string) error {
	panic("implement me")
}

func (p psqlAntibruteforceRepository) BlacklistRemove(ctx context.Context, subnet string) error {
	panic("implement me")
}

func (p psqlAntibruteforceRepository) WhitelistAdd(ctx context.Context, subnet string) error {
	panic("implement me")
}

func (p psqlAntibruteforceRepository) WhitelistRemove(ctx context.Context, subnet string) error {
	panic("implement me")
}
