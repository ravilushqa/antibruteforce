package grpc

import (
	"context"
	"gitlab.com/otus_golang/antibruteforce/internal/antibruteforce"
	apipb "gitlab.com/otus_golang/antibruteforce/internal/antibruteforce/delivery/grpc/api"
	"go.uber.org/zap"
)

type server struct {
	usecase antibruteforce.Usecase
	logger  *zap.Logger
}

func NewServer(usecase antibruteforce.Usecase, logger *zap.Logger) *server {
	return &server{usecase: usecase, logger: logger}
}

func (s *server) Check(ctx context.Context, req *apipb.CheckRequest) (*apipb.CheckResponse, error) {
	err := s.usecase.Check(ctx, req.Login, req.Password, req.Ip)

	return &apipb.CheckResponse{Ok: err == nil}, err
}

func (s *server) Reset(context.Context, *apipb.ResetRequest) (*apipb.ResetResponse, error) {
	panic("implement me")
}

func (s *server) BlacklistAdd(context.Context, *apipb.BlacklistAddRequest) (*apipb.BlacklistAddResponse, error) {
	panic("implement me")
}

func (s *server) BlacklistRemove(context.Context, *apipb.BlacklistRemoveRequest) (*apipb.BlacklistRemoveResponse, error) {
	panic("implement me")
}

func (s *server) WhitelistAdd(context.Context, *apipb.WhitelistAddRequest) (*apipb.WhitelistAddResponse, error) {
	panic("implement me")
}

func (s *server) WhitelistRemove(context.Context, *apipb.WhitelistRemoveRequest) (*apipb.WhitelistRemoveResponse, error) {
	panic("implement me")
}
