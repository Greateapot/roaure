package server

import (
	"context"

	"github.com/Greateapot/roaure/internal/database"
	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"github.com/Greateapot/roaure/internal/validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	maxDownloadThreshold = 128 * database.MBit
	maxBadCountLimit     = 10
)

// UpdateMonitorConf implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) UpdateMonitorConf(
	ctx context.Context,
	request *roaurev1.UpdateMonitorConfRequest,
) (*emptypb.Empty, error) {
	var parsedRequest updateMonitorConfRequest
	if err := parsedRequest.parse(request); err != nil {
		return nil, err
	}
	return s.updateMonitorConf(ctx, &parsedRequest)
}

func (s *roaureServiceServer) updateMonitorConf(
	_ context.Context,
	request *updateMonitorConfRequest,
) (*emptypb.Empty, error) {
	oldDownloadThreshold := s.config.MonitorConf.DownloadThreshold
	oldPollInterval := s.config.MonitorConf.PollInterval
	oldBadCountLimit := s.config.MonitorConf.BadCountLimit

	s.config.MonitorConf.DownloadThreshold = database.DataSize(request.monitorConf.DownloadThreshold)
	s.config.MonitorConf.PollInterval = &database.Time{
		Hours:   request.monitorConf.PollInterval.Hours,
		Minutes: request.monitorConf.PollInterval.Minutes,
	}
	s.config.MonitorConf.BadCountLimit = request.monitorConf.BadCountLimit

	if err := s.db.DumpConfig(s.config); err != nil {
		// Не удалось сохранить изменения, откат
		s.config.MonitorConf.DownloadThreshold = oldDownloadThreshold
		s.config.MonitorConf.PollInterval = oldPollInterval
		s.config.MonitorConf.BadCountLimit = oldBadCountLimit
		grpclog.Errorln(err)
		return nil, status.Errorf(codes.Internal, "unnable to save changes: %s", err.Error())
	}

	// Restart on PollInterval changed
	if oldPollInterval.Hours != s.config.MonitorConf.PollInterval.Hours ||
		oldPollInterval.Minutes != s.config.MonitorConf.PollInterval.Minutes {
		s.monitor.Stop()
		s.monitor.Start()
	}

	return &emptypb.Empty{}, nil
}

type updateMonitorConfRequest struct {
	monitorConf *roaurev1.MonitorConf
}

func (r *updateMonitorConfRequest) parse(request *roaurev1.UpdateMonitorConfRequest) error {
	var v validation.MessageValidator

	if request.GetMonitorConf() == nil {
		v.AddFieldViolation("monitor_conf", "required field")
	} else {
		// download_threshold = 1
		if request.GetMonitorConf().GetDownloadThreshold() == 0 {
			v.AddFieldViolation("monitor_conf.download_threshold", "required field")
		} else if request.GetMonitorConf().GetDownloadThreshold() > float64(maxDownloadThreshold) {
			v.AddFieldViolation("monitor_conf.download_threshold", "should be <= %s", maxDownloadThreshold.String())
		}
		// poll_interval = 2
		if request.GetMonitorConf().GetPollInterval() == nil {
			v.AddFieldViolation("monitor_conf.poll_interval", "required field")
		} else if request.GetMonitorConf().GetPollInterval().GetHours() > 23 {
			v.AddFieldViolation("monitor_conf.poll_interval.hours", "should be <= 23")
		} else if request.GetMonitorConf().GetPollInterval().GetMinutes() > 59 {
			v.AddFieldViolation("monitor_conf.poll_interval.minutes", "should be <= 59")
		}
		// bad_count_limit = 3
		if request.GetMonitorConf().GetBadCountLimit() == 0 {
			v.AddFieldViolation("monitor_conf.bad_count_limit", "required field")
		} else if request.GetMonitorConf().GetBadCountLimit() > maxBadCountLimit {
			v.AddFieldViolation("monitor_conf.download_threshold", "should be <= 10")
		}

		r.monitorConf = request.GetMonitorConf()
	}

	return v.Err()
}
