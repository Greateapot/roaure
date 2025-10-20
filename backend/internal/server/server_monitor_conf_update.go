package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UpdateMonitorConf implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) UpdateMonitorConf(ctx context.Context, request *roaurev1.UpdateMonitorConfRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
