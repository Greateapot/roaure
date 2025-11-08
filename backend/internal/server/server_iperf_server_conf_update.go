package server

import (
	"context"
	"regexp"

	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"github.com/Greateapot/roaure/internal/validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	maxPortNumber = 65535
)

// UpdateIperfServerConf implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) UpdateIperfServerConf(
	ctx context.Context,
	request *roaurev1.UpdateIperfServerConfRequest,
) (*emptypb.Empty, error) {
	var parsedRequest updateIperfServerConfRequest
	if err := parsedRequest.parse(request); err != nil {
		return nil, err
	}
	return s.updateIperfServerConf(ctx, &parsedRequest)
}

func (s *roaureServiceServer) updateIperfServerConf(
	_ context.Context,
	request *updateIperfServerConfRequest,
) (*emptypb.Empty, error) {
	oldHost := s.config.IperfServerConf.Host
	oldPort := s.config.IperfServerConf.Port

	s.config.IperfServerConf.Host = request.iperfServerConf.Host
	s.config.IperfServerConf.Port = request.iperfServerConf.Port

	if err := s.db.DumpConfig(s.config); err != nil {
		// Не удалось сохранить изменения, откат
		s.config.IperfServerConf.Host = oldHost
		s.config.IperfServerConf.Port = oldPort
		grpclog.Errorln(err)
		return nil, status.Errorf(codes.Internal, "unnable to save changes: %s", err.Error())
	}

	s.monitor.SpeedtestClient.SetupClient()

	return &emptypb.Empty{}, nil
}

type updateIperfServerConfRequest struct {
	iperfServerConf *roaurev1.IperfServerConf
}

func (r *updateIperfServerConfRequest) parse(request *roaurev1.UpdateIperfServerConfRequest) error {
	var v validation.MessageValidator

	if request.GetIperfServerConf() == nil {
		v.AddFieldViolation("server_info", "required field")
	} else {
		// host = 1
		if len(request.GetIperfServerConf().GetHost()) == 0 {
			v.AddFieldViolation("server_info.host", "required field")
		} else if matched, _ := regexp.MatchString(hostnamePattern, request.GetIperfServerConf().GetHost()); !matched {
			v.AddFieldViolation("server_info.host", "invalid hostname format")
		}
		// port = 2
		if request.GetIperfServerConf().GetPort() == 0 {
			v.AddFieldViolation("router_info.port", "required field")
		} else if request.GetIperfServerConf().GetPort() > maxPortNumber {
			v.AddFieldViolation("router_info.port", "should be <= 65535")
		}

		r.iperfServerConf = request.GetIperfServerConf()
	}

	return v.Err()
}
