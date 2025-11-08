package server

import (
	"context"
	"slices"

	"github.com/Greateapot/roaure/internal/database"
	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"github.com/Greateapot/roaure/internal/validation"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteSchedule implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) DeleteSchedule(
	ctx context.Context,
	request *roaurev1.DeleteScheduleRequest,
) (*emptypb.Empty, error) {
	var parsedRequest deleteScheduleRequest
	if err := parsedRequest.parse(request); err != nil {
		return nil, err
	}
	return s.deleteSchedule(ctx, &parsedRequest)

}

func (s *roaureServiceServer) deleteSchedule(
	_ context.Context,
	request *deleteScheduleRequest,
) (*emptypb.Empty, error) {
	scheduleIndex := slices.IndexFunc(s.config.MonitorConf.Schedules, func(schedule *database.Schedule) bool {
		return schedule.ID == request.id
	})
	if scheduleIndex == -1 {
		return nil, status.Errorf(codes.NotFound, "schedule with id %s not found", request.id.String())
	}

	schedule := s.config.MonitorConf.Schedules[scheduleIndex]

	s.config.MonitorConf.Schedules = slices.Delete(s.config.MonitorConf.Schedules, scheduleIndex, scheduleIndex+1)
	if err := s.db.DumpConfig(s.config); err != nil {
		// Не удалось сохранить изменения, откат
		s.config.MonitorConf.Schedules = append(s.config.MonitorConf.Schedules, schedule)
		grpclog.Errorln(err)
		return nil, status.Errorf(codes.Internal, "unnable to save changes: %s", err.Error())
	}

	return &emptypb.Empty{}, nil
}

type deleteScheduleRequest struct {
	id uuid.UUID
}

func (r *deleteScheduleRequest) parse(request *roaurev1.DeleteScheduleRequest) error {
	var v validation.MessageValidator

	if len(request.GetId()) == 0 {
		v.AddFieldViolation("id", "required field")
	} else {
		id, err := uuid.Parse(request.GetId())
		if err != nil {
			v.AddFieldViolation("id", err.Error())
		}

		r.id = id
	}

	return v.Err()
}
