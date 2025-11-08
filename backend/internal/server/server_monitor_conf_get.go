package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetMonitorConf implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) GetMonitorConf(ctx context.Context, request *emptypb.Empty) (*roaurev1.MonitorConf, error) {
	return &roaurev1.MonitorConf{
		DownloadThreshold: s.config.MonitorConf.DownloadThreshold.Float(),
		PollInterval: &roaurev1.Time{
			Hours:   int32(s.config.MonitorConf.PollInterval.Hours),
			Minutes: int32(s.config.MonitorConf.PollInterval.Minutes),
		},
		BadCountLimit: s.config.MonitorConf.BadCountLimit,
	}, nil
}
