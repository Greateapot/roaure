package server

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/Greateapot/roaure/internal/database"
	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"github.com/Greateapot/roaure/internal/validation"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

const (
	maxTitleLength = 64
)

// CreateSchedule implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) CreateSchedule(
	ctx context.Context,
	request *roaurev1.CreateScheduleRequest,
) (*roaurev1.Schedule, error) {
	var parsedRequest createScheduleRequest
	if err := parsedRequest.parse(request); err != nil {
		return nil, err
	}
	return s.createSchedule(ctx, &parsedRequest)
}

func (s *roaureServiceServer) createSchedule(
	_ context.Context,
	request *createScheduleRequest,
) (*roaurev1.Schedule, error) {
	id, err := uuid.NewV7()
	if err != nil {
		grpclog.Errorln(err)
		return nil, status.Error(codes.Internal, "unnable to generate new UUIDv7")
	}

	schedule := &database.Schedule{
		ID:    id,
		Title: request.schedule.Title,
		StartsAt: &database.Time{
			Hours:   request.schedule.StartsAt.Hours,
			Minutes: request.schedule.StartsAt.Minutes,
		},
		EndsAt: &database.Time{
			Hours:   request.schedule.EndsAt.Hours,
			Minutes: request.schedule.EndsAt.Minutes,
		},
		Weekdays: request.weekdays,
		Enabled:  request.schedule.Enabled,
	}

	s.config.MonitorConf.Schedules = append(s.config.MonitorConf.Schedules, schedule)
	if err := s.db.DumpConfig(s.config); err != nil {
		// Не удалось сохранить изменения, откат
		s.config.MonitorConf.Schedules = slices.Delete(
			s.config.MonitorConf.Schedules,
			len(s.config.MonitorConf.Schedules)-1,
			len(s.config.MonitorConf.Schedules),
		)
		grpclog.Errorln(err)
		return nil, status.Errorf(codes.Internal, "unnable to save changes: %s", err.Error())
	}

	request.schedule.Id = id.String()

	return request.schedule, nil
}

type createScheduleRequest struct {
	schedule *roaurev1.Schedule
	weekdays []time.Weekday
}

func (r *createScheduleRequest) parse(request *roaurev1.CreateScheduleRequest) error {
	var v validation.MessageValidator

	if request.GetSchedule() == nil {
		v.AddFieldViolation("schedule", "required field")
	} else {
		// title = 2
		if len(request.GetSchedule().GetTitle()) == 0 {
			v.AddFieldViolation("schedule.title", "required field")
		} else if len(request.GetSchedule().GetTitle()) > maxTitleLength {
			v.AddFieldViolation("schedule.title", "should be <= 63 characters")
		}
		// start_at = 3
		if request.GetSchedule().GetStartsAt() == nil {
			v.AddFieldViolation("schedule.starts_at", "required field")
		} else {
			if request.GetSchedule().GetStartsAt().GetHours() > 23 {
				v.AddFieldViolation("schedule.starts_at.hours", "should be <= 23")
			}
			if request.GetSchedule().GetStartsAt().GetMinutes() > 59 {
				v.AddFieldViolation("schedule.starts_at.minutes", "should be <= 59")
			}
		}
		// ends_at = 4
		if request.GetSchedule().GetEndsAt() == nil {
			v.AddFieldViolation("schedule.ends_at", "required field")
		} else {
			if request.GetSchedule().GetEndsAt().GetHours() > 23 {
				v.AddFieldViolation("schedule.ends_at.hours", "should be <= 23")
			}
			if request.GetSchedule().GetEndsAt().GetMinutes() > 59 {
				v.AddFieldViolation("schedule.ends_at.minutes", "should be <= 59")
			}
		}
		// weekdays = 5
		r.weekdays = make([]time.Weekday, len(request.GetSchedule().GetWeekdays()))
		for i, weekday := range request.GetSchedule().GetWeekdays() {
			if w, err := ConvertWeekdayProtoToTime(weekday); err != nil {
				v.AddFieldViolation(fmt.Sprintf("schedule.weekdays[%d]", i), err.Error())
			} else {
				r.weekdays[i] = w
			}
		}

		r.schedule = request.GetSchedule()
	}

	return v.Err()
}
