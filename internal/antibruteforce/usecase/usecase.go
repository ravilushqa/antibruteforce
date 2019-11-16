package usecase

import (
	"context"
	"gitlab.com/otus_golang/antibruteforce/internal/antibruteforce"
	"gitlab.com/otus_golang/antibruteforce/internal/bucket"
	"go.uber.org/zap"
)

type antibruteforceUsecase struct {
	antibruteforceRepo *antibruteforce.Repository
	bucketRepo         *bucket.Repository
	logger             *zap.Logger
}

func NewAntibruteforceUsecase(antibruteforceRepo *antibruteforce.Repository, bucketRepo *bucket.Repository, logger *zap.Logger) *antibruteforceUsecase {
	return &antibruteforceUsecase{antibruteforceRepo: antibruteforceRepo, bucketRepo: bucketRepo, logger: logger}
}

func (a antibruteforceUsecase) Check(ctx context.Context, login string, password string, ip string) error {
	panic("implement me")
}

func (a antibruteforceUsecase) Reset(ctx context.Context, login string, ip string) error {
	panic("implement me")
}

func (a antibruteforceUsecase) BlacklistAdd(ctx context.Context, subnet string) error {
	panic("implement me")
}

func (a antibruteforceUsecase) BlacklistRemove(ctx context.Context, subnet string) error {
	panic("implement me")
}

func (a antibruteforceUsecase) WhitelistAdd(ctx context.Context, subnet string) error {
	panic("implement me")
}

func (a antibruteforceUsecase) WhitelistRemove(ctx context.Context, subnet string) error {
	panic("implement me")
}
