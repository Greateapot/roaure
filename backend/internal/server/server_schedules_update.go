package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateSchedule implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) UpdateSchedule(ctx context.Context, request *roaurev1.UpdateScheduleRequest) (*roaurev1.Schedule, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
