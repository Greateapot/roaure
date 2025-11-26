package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Reboot implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) Reboot(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	if err := s.monitor.Reboot(); err != nil {
		return nil, status.Errorf(codes.Internal, "unnable to reboot router: %s", err.Error())
	} else {
		return &emptypb.Empty{}, nil
	}
}
