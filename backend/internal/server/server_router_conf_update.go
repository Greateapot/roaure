package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UpdateRouterConf implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) UpdateRouterConf(ctx context.Context, request *roaurev1.UpdateRouterConfRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
