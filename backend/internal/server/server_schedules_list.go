package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ListSchedules implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) ListSchedules(
	ctx context.Context,
	request *emptypb.Empty,
) (*roaurev1.ListSchedulesResponse, error) {
	response := roaurev1.ListSchedulesResponse{
		Schedules: make([]*roaurev1.Schedule, 0, len(s.config.MonitorConf.Schedules)),
	}

	for _, schedule := range s.config.MonitorConf.Schedules {
		response.Schedules = append(response.Schedules, ConvertScheduleRowToProto(schedule))
	}

	return &response, nil
}
