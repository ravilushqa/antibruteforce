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

func (s *server) Reset(ctx context.Context, req *apipb.ResetRequest) (*apipb.ResetResponse, error) {
	err := s.usecase.Reset(ctx, req.Login, req.Ip)

	return &apipb.ResetResponse{Ok: err == nil}, err
}

func (s *server) BlacklistAdd(ctx context.Context, req *apipb.BlacklistAddRequest) (*apipb.BlacklistAddResponse, error) {
	err := s.usecase.BlacklistAdd(ctx, req.Subnet)
	return &apipb.BlacklistAddResponse{Ok: err == nil}, err
}

func (s *server) BlacklistRemove(ctx context.Context, req *apipb.BlacklistRemoveRequest) (*apipb.BlacklistRemoveResponse, error) {
	err := s.usecase.BlacklistRemove(ctx, req.Subnet)
	return &apipb.BlacklistRemoveResponse{Ok: err == nil}, err
}

func (s *server) WhitelistAdd(ctx context.Context, req *apipb.WhitelistAddRequest) (*apipb.WhitelistAddResponse, error) {
	err := s.usecase.WhitelistAdd(ctx, req.Subnet)
	return &apipb.WhitelistAddResponse{Ok: err == nil}, err
}

func (s *server) WhitelistRemove(ctx context.Context, req *apipb.WhitelistRemoveRequest) (*apipb.WhitelistRemoveResponse, error) {
	err := s.usecase.WhitelistRemove(ctx, req.Subnet)
	return &apipb.WhitelistRemoveResponse{Ok: err == nil}, err
}
