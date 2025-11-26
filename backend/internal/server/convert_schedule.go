package server

import (
	"github.com/Greateapot/roaure/internal/database"
	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
)

func ConvertScheduleRowToProto(schedule *database.Schedule) *roaurev1.Schedule {
	weekdays := make([]roaurev1.Weekday, 0, len(schedule.Weekdays))

	for _, weekday := range schedule.Weekdays {
		weekdays = append(weekdays, ConvertWeekdayTimeToProto(weekday))
	}

	return &roaurev1.Schedule{
		Id:    schedule.ID.String(),
		Title: schedule.Title,
		StartsAt: &roaurev1.Time{
			Hours:   uint32(schedule.StartsAt.Hours),
			Minutes: uint32(schedule.StartsAt.Minutes),
		},
		EndsAt: &roaurev1.Time{
			Hours:   uint32(schedule.EndsAt.Hours),
			Minutes: uint32(schedule.EndsAt.Minutes),
		},
		Weekdays: weekdays,
		Enabled:  schedule.Enabled,
	}
}
