package server

import (
	"time"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

const (
	maxPollInterval     = 5 * 60 // 05 мин
	defaultPollInterval = 30     // 30 сек
)

// GetMetrics implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) GetMetrics(
	request *roaurev1.GetMetricsRequest,
	server grpc.ServerStreamingServer[roaurev1.Metrics],
) error {
	switch {
	case request.GetPollInterval() > maxPollInterval:
		return status.Errorf(codes.InvalidArgument, "max max poll interval is %d seconds", maxPollInterval)
	case request.GetPollInterval() == 0:
		request.PollInterval = defaultPollInterval
	}

	pollInterval := time.Duration(request.PollInterval) * time.Second

	for {
		if err := server.Send(&roaurev1.Metrics{
			DownloadSpeed:  float64(s.monitor.DownloadSpeed),
			RebootRequired: s.monitor.RebootRequired,
			BadCount:       uint32(s.monitor.BadCount),
			MonitorRunning: s.monitor.Running,
		}); err != nil {
			if s, ok := status.FromError(err); !ok || s.Code() != codes.Canceled {
				grpclog.Errorln(err)
			}
			return nil
		}

		// sleep
		<-time.After(pollInterval)
	}
}
