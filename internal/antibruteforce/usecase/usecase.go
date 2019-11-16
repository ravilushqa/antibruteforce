package usecase

import (
	"context"
	"gitlab.com/otus_golang/antibruteforce/config"
	"gitlab.com/otus_golang/antibruteforce/internal/antibruteforce"
	"gitlab.com/otus_golang/antibruteforce/internal/bucket"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"time"
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

	var g errgroup.Group

	adds := []struct {
		key      string
		capacity uint
	}{
		{
			login,
			a.config.BucketLoginCapacity,
		},
		{
			password,
			a.config.BucketPasswordCapacity,
		},
		{
			ip,
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
	panic("implement me")
}

func (a *antibruteforceUsecase) BlacklistAdd(ctx context.Context, subnet string) error {
	panic("implement me")
}

func (a *antibruteforceUsecase) BlacklistRemove(ctx context.Context, subnet string) error {
	panic("implement me")
}

func (a *antibruteforceUsecase) WhitelistAdd(ctx context.Context, subnet string) error {
	panic("implement me")
}

func (a *antibruteforceUsecase) WhitelistRemove(ctx context.Context, subnet string) error {
	panic("implement me")
}
