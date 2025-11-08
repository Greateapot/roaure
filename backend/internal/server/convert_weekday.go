package server

import (
	"fmt"
	"time"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
)

func ConvertWeekdayTimeToProto(weekday time.Weekday) roaurev1.Weekday {
	switch weekday {
	case time.Sunday:
		return roaurev1.Weekday_WEEKDAY_SUNDAY
	case time.Monday:
		return roaurev1.Weekday_WEEKDAY_MONDAY
	case time.Tuesday:
		return roaurev1.Weekday_WEEKDAY_TUESDAY
	case time.Wednesday:
		return roaurev1.Weekday_WEEKDAY_WEDNESDAY
	case time.Thursday:
		return roaurev1.Weekday_WEEKDAY_THURSDAY
	case time.Friday:
		return roaurev1.Weekday_WEEKDAY_FRIDAY
	case time.Saturday:
		return roaurev1.Weekday_WEEKDAY_SATURDAY
	default:
		return roaurev1.Weekday_WEEKDAY_UNSPECIFIED
	}
}

func ConvertWeekdayProtoToTime(weekday roaurev1.Weekday) (time.Weekday, error) {
	switch weekday {
	case roaurev1.Weekday_WEEKDAY_SUNDAY:
		return time.Sunday, nil
	case roaurev1.Weekday_WEEKDAY_MONDAY:
		return time.Monday, nil
	case roaurev1.Weekday_WEEKDAY_TUESDAY:
		return time.Tuesday, nil
	case roaurev1.Weekday_WEEKDAY_WEDNESDAY:
		return time.Wednesday, nil
	case roaurev1.Weekday_WEEKDAY_THURSDAY:
		return time.Thursday, nil
	case roaurev1.Weekday_WEEKDAY_FRIDAY:
		return time.Friday, nil
	case roaurev1.Weekday_WEEKDAY_SATURDAY:
		return time.Saturday, nil
	default: // roaurev1.Weekday_WEEKDAY_UNSPECIFIED & etc
		return -1, fmt.Errorf("unexpected roaurev1.Weekday: %#v", weekday)
	}
}
