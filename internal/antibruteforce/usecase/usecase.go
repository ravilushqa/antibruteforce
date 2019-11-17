package usecase

import (
	"context"
	"gitlab.com/otus_golang/antibruteforce/config"
	"gitlab.com/otus_golang/antibruteforce/internal/antibruteforce"
	"gitlab.com/otus_golang/antibruteforce/internal/antibruteforce/consts"
	"gitlab.com/otus_golang/antibruteforce/internal/antibruteforce/errors"
	"gitlab.com/otus_golang/antibruteforce/internal/bucket"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net"
	"time"
)

const (
	LOGIN_PREFIX    = "login_"
	PASSWORD_PREFIX = "password_"
	IP_PREFIX       = "ip"
)

type antibruteforceUsecase struct {
	antibruteforceRepo antibruteforce.Repository
	bucketRepo         bucket.Repository
	logger             *zap.Logger
	config             *config.Config
}

func NewAntibruteforceUsecase(antibruteforceRepo antibruteforce.Repository, bucketRepo bucket.Repository, logger *zap.Logger, config *config.Config) *antibruteforceUsecase {
	return &antibruteforceUsecase{antibruteforceRepo: antibruteforceRepo, bucketRepo: bucketRepo, logger: logger, config: config}
}

func (a *antibruteforceUsecase) Check(ctx context.Context, login string, password string, ip string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	list, err := a.antibruteforceRepo.FindIpInList(ctx, ip)
	if err != nil {
		return err
	}
	switch list {
	case consts.Whitelist:
		return nil
	case consts.Blacklist:
		return errors.ErrIpInBlackList
	}

	var g errgroup.Group

	adds := []struct {
		key      string
		capacity uint
	}{
		{
			LOGIN_PREFIX + login,
			a.config.BucketLoginCapacity,
		},
		{
			PASSWORD_PREFIX + password,
			a.config.BucketPasswordCapacity,
		},
		{
			IP_PREFIX + ip,
			a.config.BucketIpCapacity,
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

func (a *antibruteforceUsecase) Reset(ctx context.Context, login string, ip string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	return a.bucketRepo.Reset(ctx, []string{LOGIN_PREFIX + login, IP_PREFIX + ip})
}

func (a *antibruteforceUsecase) BlacklistAdd(ctx context.Context, subnet string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	_, _, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}
	return a.antibruteforceRepo.BlacklistAdd(ctx, subnet)
}

func (a *antibruteforceUsecase) BlacklistRemove(ctx context.Context, subnet string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	_, _, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}
	return a.antibruteforceRepo.BlacklistRemove(ctx, subnet)
}

func (a *antibruteforceUsecase) WhitelistAdd(ctx context.Context, subnet string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	_, _, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}
	return a.antibruteforceRepo.WhitelistAdd(ctx, subnet)
}

func (a *antibruteforceUsecase) WhitelistRemove(ctx context.Context, subnet string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.ContextTimeout)*time.Millisecond)
	defer cancel()

	_, _, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}
	return a.antibruteforceRepo.WhitelistRemove(ctx, subnet)
}
