package server

import (
	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetMetrics implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) GetMetrics(request *roaurev1.GetMetricsRequest, server grpc.ServerStreamingServer[roaurev1.Metrics]) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}
