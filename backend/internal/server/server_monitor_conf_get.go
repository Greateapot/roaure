package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetMonitorConf implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) GetMonitorConf(ctx context.Context, request *emptypb.Empty) (*roaurev1.MonitorConf, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
