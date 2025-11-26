package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetIperfServerConf implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) GetIperfServerConf(context.Context, *emptypb.Empty) (*roaurev1.IperfServerConf, error) {
	return &roaurev1.IperfServerConf{
		Host: s.config.IperfServerConf.Host,
		Port: int32(s.config.IperfServerConf.Port),
	}, nil
}
