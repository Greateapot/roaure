package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetMonitorConf implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) GetMonitorConf(context.Context, *emptypb.Empty) (*roaurev1.MonitorConf, error) {
	return &roaurev1.MonitorConf{
		DownloadThreshold: float64(s.config.MonitorConf.DownloadThreshold),
		PollInterval: &roaurev1.Time{
			Hours:   uint32(s.config.MonitorConf.PollInterval.Hours),
			Minutes: uint32(s.config.MonitorConf.PollInterval.Minutes),
		},
		BadCountLimit: uint32(s.config.MonitorConf.BadCountLimit),
	}, nil
}
