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

// UpdateSchedule implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) UpdateSchedule(
	ctx context.Context,
	request *roaurev1.UpdateScheduleRequest,
) (*roaurev1.Schedule, error) {
	var parsedRequest updateScheduleRequest
	if err := parsedRequest.parse(request); err != nil {
		return nil, err
	}
	return s.updateSchedule(ctx, &parsedRequest)
}

func (s *roaureServiceServer) updateSchedule(
	_ context.Context,
	request *updateScheduleRequest,
) (*roaurev1.Schedule, error) {
	scheduleIndex := slices.IndexFunc(s.config.MonitorConf.Schedules, func(schedule *database.Schedule) bool {
		return schedule.ID == request.id
	})
	if scheduleIndex == -1 {
		return nil, status.Errorf(codes.NotFound, "schedule with id %s not found", request.id.String())
	}

	oldSchedule := s.config.MonitorConf.Schedules[scheduleIndex]

	newSchedule := &database.Schedule{
		ID:    request.id,
		Title: request.schedule.Title,
		StartsAt: &database.Time{
			Hours:   uint8(request.schedule.StartsAt.Hours),
			Minutes: uint8(request.schedule.StartsAt.Minutes),
		},
		EndsAt: &database.Time{
			Hours:   uint8(request.schedule.EndsAt.Hours),
			Minutes: uint8(request.schedule.EndsAt.Minutes),
		},
		Weekdays: request.weekdays,
		Enabled:  request.schedule.Enabled,
	}

	s.config.MonitorConf.Schedules[scheduleIndex] = newSchedule
	if err := s.Database.DumpConfig(s.config); err != nil {
		// Не удалось сохранить изменения, откат
		s.config.MonitorConf.Schedules[scheduleIndex] = oldSchedule
		grpclog.Errorln(err)
		return nil, status.Errorf(codes.Internal, "unnable to save changes: %s", err.Error())
	}

	request.schedule.Id = request.id.String()

	return request.schedule, nil
}

type updateScheduleRequest struct {
	id       uuid.UUID
	schedule *roaurev1.Schedule
	weekdays []time.Weekday
}

func (r *updateScheduleRequest) parse(request *roaurev1.UpdateScheduleRequest) error {
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
