package usecase

import (
	"context"
	"github.com/ravilushqa/antibruteforce/config"
	"github.com/ravilushqa/antibruteforce/internal/antibruteforce"
	"github.com/ravilushqa/antibruteforce/internal/antibruteforce/consts"
	"github.com/ravilushqa/antibruteforce/internal/antibruteforce/errors"
	"github.com/ravilushqa/antibruteforce/internal/bucket"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net"
	"time"
)

// AntibruteforceUsecase implements antibruteforce usecase
type AntibruteforceUsecase struct {
	antibruteforceRepo antibruteforce.Repository
	bucketRepo         bucket.Repository
	logger             *zap.Logger
	config             *config.Config
}

// NewAntibruteforceUsecase Constructor for AntibruteforceUsecase
func NewAntibruteforceUsecase(antibruteforceRepo antibruteforce.Repository, bucketRepo bucket.Repository, logger *zap.Logger, config *config.Config) *AntibruteforceUsecase {
	return &AntibruteforceUsecase{antibruteforceRepo: antibruteforceRepo, bucketRepo: bucketRepo, logger: logger, config: config}
}

// Check method Checking that requester is able to send request.
func (a *AntibruteforceUsecase) Check(ctx context.Context, login string, password string, ip string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	if login == "" {
		return errors.ErrLoginRequired
	}

	if password == "" {
		return errors.ErrPasswordRequired
	}

	if net.ParseIP(ip) == nil {
		return errors.ErrWrongIP
	}

	list, err := a.antibruteforceRepo.FindIPInList(ctx, ip)
	if err != nil {
		return err
	}
	switch list {
	case consts.Whitelist:
		return nil
	case consts.Blacklist:
		return errors.ErrIPInBlackList
	}

	var g errgroup.Group

	adds := []struct {
		key      string
		capacity uint
	}{
		{
			consts.LoginPrefix + login,
			a.config.BucketLoginCapacity,
		},
		{
			consts.PasswordPrefix + password,
			a.config.BucketPasswordCapacity,
		},
		{
			consts.IPPrefix + ip,
			a.config.BucketIPCapacity,
		},
	}
	for _, add := range adds {
		add := add // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			return a.bucketRepo.Add(ctx, add.key, add.capacity, time.Duration(a.config.BucketRate)*time.Second)
		})
	}

	return g.Wait()
}

// Reset method is resets visitor buckets
func (a *AntibruteforceUsecase) Reset(ctx context.Context, login string, ip string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	if login == "" {
		return errors.ErrLoginRequired
	}

	if net.ParseIP(ip) == nil {
		return errors.ErrWrongIP
	}

	return a.bucketRepo.Reset(ctx, []string{consts.LoginPrefix + login, consts.IPPrefix + ip})
}

// BlacklistAdd method adding subnet to blacklist
func (a *AntibruteforceUsecase) BlacklistAdd(ctx context.Context, subnet string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	_, _, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}
	return a.antibruteforceRepo.BlacklistAdd(ctx, subnet)
}

// BlacklistRemove method removing subnet from blacklist
func (a *AntibruteforceUsecase) BlacklistRemove(ctx context.Context, subnet string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	_, _, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}
	return a.antibruteforceRepo.BlacklistRemove(ctx, subnet)
}

// WhitelistAdd method adding subnet to whitelist
func (a *AntibruteforceUsecase) WhitelistAdd(ctx context.Context, subnet string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	_, _, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}
	return a.antibruteforceRepo.WhitelistAdd(ctx, subnet)
}

// WhitelistRemove method removing subnet from whitelist
func (a *AntibruteforceUsecase) WhitelistRemove(ctx context.Context, subnet string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	_, _, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}
	return a.antibruteforceRepo.WhitelistRemove(ctx, subnet)
}
