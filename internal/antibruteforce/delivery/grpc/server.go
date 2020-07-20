package grpc

import (
	"context"
	"github.com/ravilushqa/antibruteforce/internal/antibruteforce"
	apipb "github.com/ravilushqa/antibruteforce/internal/antibruteforce/delivery/grpc/api"
	"go.uber.org/zap"
)

// Server is a struct that implements antibruteforce grpc server
type Server struct {
	usecase antibruteforce.Usecase
	logger  *zap.Logger
}

// NewServer constructor for grpc Server
func NewServer(usecase antibruteforce.Usecase, logger *zap.Logger) *Server {
	return &Server{usecase: usecase, logger: logger}
}

// Check method is checking that visitor is able to make request
func (s *Server) Check(ctx context.Context, req *apipb.CheckRequest) (*apipb.CheckResponse, error) {
	err := s.usecase.Check(ctx, req.Login, req.Password, req.Ip)

	return &apipb.CheckResponse{Ok: err == nil}, err
}

// Reset method is resets visitor buckets
func (s *Server) Reset(ctx context.Context, req *apipb.ResetRequest) (*apipb.ResetResponse, error) {
	err := s.usecase.Reset(ctx, req.Login, req.Ip)

	return &apipb.ResetResponse{Ok: err == nil}, err
}

// BlacklistAdd method adding subnet to blacklist
func (s *Server) BlacklistAdd(ctx context.Context, req *apipb.BlacklistAddRequest) (*apipb.BlacklistAddResponse, error) {
	err := s.usecase.BlacklistAdd(ctx, req.Subnet)
	return &apipb.BlacklistAddResponse{Ok: err == nil}, err
}

// BlacklistRemove method removing subnet from blacklist
func (s *Server) BlacklistRemove(ctx context.Context, req *apipb.BlacklistRemoveRequest) (*apipb.BlacklistRemoveResponse, error) {
	err := s.usecase.BlacklistRemove(ctx, req.Subnet)
	return &apipb.BlacklistRemoveResponse{Ok: err == nil}, err
}

// WhitelistAdd method adding subnet to whitelist
func (s *Server) WhitelistAdd(ctx context.Context, req *apipb.WhitelistAddRequest) (*apipb.WhitelistAddResponse, error) {
	err := s.usecase.WhitelistAdd(ctx, req.Subnet)
	return &apipb.WhitelistAddResponse{Ok: err == nil}, err
}

// WhitelistRemove method removing subnet from whitelist
func (s *Server) WhitelistRemove(ctx context.Context, req *apipb.WhitelistRemoveRequest) (*apipb.WhitelistRemoveResponse, error) {
	err := s.usecase.WhitelistRemove(ctx, req.Subnet)
	return &apipb.WhitelistRemoveResponse{Ok: err == nil}, err
}
