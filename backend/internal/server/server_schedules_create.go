package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateSchedule implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) CreateSchedule(ctx context.Context, request *roaurev1.CreateScheduleRequest) (*roaurev1.Schedule, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
