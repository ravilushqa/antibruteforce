package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/otus_golang/antibruteforce/internal/antibruteforce/consts"
	"go.uber.org/zap"
)

type psqlAntibruteforceRepository struct {
	*sqlx.DB
	logger *zap.Logger
}

func NewPsqlAntibruteforceRepository(DB *sqlx.DB, logger *zap.Logger) *psqlAntibruteforceRepository {
	return &psqlAntibruteforceRepository{DB: DB, logger: logger}
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

func (p psqlAntibruteforceRepository) FindIpInList(ctx context.Context, ip string) (string, error) {
	query := `
		SELECT distinct $2 as list FROM blacklist where $1::inet <<= subnet
		union (
			SELECT distinct $3 as list FROM whitelist where $1::inet <<= subnet
		)
	`
	sliceList := make([]string, 0, 2)

	err := p.DB.SelectContext(ctx, &sliceList, query, ip, consts.Blacklist, consts.Whitelist)

	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	switch len(sliceList) {
	case 0:
		return "", nil
	case 1:
		return sliceList[0], nil
	default:
		p.logger.Info(fmt.Sprintf("ip: %s in more than one list. lists: %v", ip, sliceList))
		return consts.Blacklist, nil
	}
}
