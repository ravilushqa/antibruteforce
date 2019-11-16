package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type psqlAntibruteforceRepository struct {
	*sqlx.DB
}

func NewPsqlAntibruteforceRepository(DB *sqlx.DB) *psqlAntibruteforceRepository {
	return &psqlAntibruteforceRepository{DB: DB}
}

func (p psqlAntibruteforceRepository) BlacklistAdd(ctx context.Context, subnet string) error {
	query := `INSERT INTO blacklist (subnet) VALUES ($1)`
	_, err := p.DB.ExecContext(ctx, query, subnet)

	return err
}

func (p psqlAntibruteforceRepository) BlacklistRemove(ctx context.Context, subnet string) error {
	query := `DELETE FROM blacklist WHERE subnet = $1`
	_, err := p.DB.ExecContext(ctx, query, subnet)

	return err
}

func (p psqlAntibruteforceRepository) WhitelistAdd(ctx context.Context, subnet string) error {
	query := `INSERT INTO whitelist (subnet) VALUES ($1)`
	_, err := p.DB.ExecContext(ctx, query, subnet)

	return err
}

func (p psqlAntibruteforceRepository) WhitelistRemove(ctx context.Context, subnet string) error {
	query := `DELETE FROM whitelist WHERE subnet = $1`
	_, err := p.DB.ExecContext(ctx, query, subnet)

	return err
}
