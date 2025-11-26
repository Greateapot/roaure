package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

// ToggleMonitor implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) ToggleMonitor(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	if s.monitor.Running {
		s.monitor.Stop()
	} else {
		s.monitor.Start()
	}

	return &emptypb.Empty{}, nil
}
