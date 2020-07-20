package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ravilushqa/antibruteforce/internal/antibruteforce/consts"
	"go.uber.org/zap"
)

// PsqlAntibruteforceRepository is Repository implementation for postgres
type PsqlAntibruteforceRepository struct {
	*sqlx.DB
	logger *zap.Logger
}

// NewPsqlAntibruteforceRepository constructor for postgres repository
func NewPsqlAntibruteforceRepository(DB *sqlx.DB, logger *zap.Logger) *PsqlAntibruteforceRepository {
	return &PsqlAntibruteforceRepository{DB: DB, logger: logger}
}

// BlacklistAdd adding subnet to blacklist
func (p PsqlAntibruteforceRepository) BlacklistAdd(ctx context.Context, subnet string) error {
	query := `INSERT INTO blacklist (subnet) VALUES ($1)`
	_, err := p.DB.ExecContext(ctx, query, subnet)

	return err
}

// BlacklistRemove removing subnet from blacklist
func (p PsqlAntibruteforceRepository) BlacklistRemove(ctx context.Context, subnet string) error {
	query := `DELETE FROM blacklist WHERE subnet = $1`
	_, err := p.DB.ExecContext(ctx, query, subnet)

	return err
}

// WhitelistAdd adding subnet to whitelist
func (p PsqlAntibruteforceRepository) WhitelistAdd(ctx context.Context, subnet string) error {
	query := `INSERT INTO whitelist (subnet) VALUES ($1)`
	_, err := p.DB.ExecContext(ctx, query, subnet)

	return err
}

// WhitelistRemove removing subnet from whitelist
func (p PsqlAntibruteforceRepository) WhitelistRemove(ctx context.Context, subnet string) error {
	query := `DELETE FROM whitelist WHERE subnet = $1`
	_, err := p.DB.ExecContext(ctx, query, subnet)

	return err
}

// FindIPInList finding IP in blacklist or whitelist, possible values: "blacklist", "whitelist", ""
// If ip find in both, returns "blacklist"
func (p PsqlAntibruteforceRepository) FindIPInList(ctx context.Context, ip string) (string, error) {
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
