package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetRouterConf implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) GetRouterConf(ctx context.Context, request *emptypb.Empty) (*roaurev1.RouterConf, error) {
	return &roaurev1.RouterConf{
		Host:     s.config.RouterConf.Host,
		Username: s.config.RouterConf.Username,
		Password: s.config.RouterConf.Password,
	}, nil
}
