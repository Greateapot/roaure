package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ListSchedules implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) ListSchedules(ctx context.Context, request *emptypb.Empty) (*roaurev1.ListSchedulesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
