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
			DownloadSpeed:  s.monitor.DownloadSpeed.Float(),
			RebootRequired: s.monitor.RebootRequired,
			BadCount:       s.monitor.BadCount,
			MonitorRunning: s.monitor.Running,
		}); err != nil {
			st, ok := status.FromError(err)
			if !ok {
				// runtime error
				panic(err)
			}

			switch st.Code() {
			case codes.Canceled, codes.Unavailable:
				// valid codes, skip
			default:
				// log it
				grpclog.Errorln(err)
			}

			return nil
		}

		// sleep
		<-time.After(pollInterval)
	}
}
