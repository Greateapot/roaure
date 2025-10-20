package server

import (
	"context"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetIperfServerConf implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) GetIperfServerConf(ctx context.Context, request *emptypb.Empty) (*roaurev1.IperfServerConf, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
