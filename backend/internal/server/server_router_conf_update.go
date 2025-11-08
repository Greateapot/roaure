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
	hostnamePattern   = `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`
	maxUsernameLength = 256
	maxPasswordLength = 256
)

// UpdateRouterConf implements roaurev1.RoaureServiceServer.
func (s *roaureServiceServer) UpdateRouterConf(
	ctx context.Context,
	request *roaurev1.UpdateRouterConfRequest,
) (*emptypb.Empty, error) {
	var parsedRequest updateRouterConfRequest
	if err := parsedRequest.parse(request); err != nil {
		return nil, err
	}
	return s.updateRouterConf(ctx, &parsedRequest)
}

func (s *roaureServiceServer) updateRouterConf(
	_ context.Context,
	request *updateRouterConfRequest,
) (*emptypb.Empty, error) {
	oldHost := s.config.RouterConf.Host
	oldUsername := s.config.RouterConf.Username
	oldPassword := s.config.RouterConf.Password

	s.config.RouterConf.Host = request.routerConf.Host
	s.config.RouterConf.Username = request.routerConf.Username
	s.config.RouterConf.Password = request.routerConf.Password

	if err := s.db.DumpConfig(s.config); err != nil {
		// Не удалось сохранить изменения, откат
		s.config.RouterConf.Host = oldHost
		s.config.RouterConf.Username = oldUsername
		s.config.RouterConf.Password = oldPassword
		grpclog.Errorln(err)
		return nil, status.Errorf(codes.Internal, "unnable to save changes: %s", err.Error())
	}

	return &emptypb.Empty{}, nil
}

type updateRouterConfRequest struct {
	routerConf *roaurev1.RouterConf
}

func (r *updateRouterConfRequest) parse(request *roaurev1.UpdateRouterConfRequest) error {
	var v validation.MessageValidator

	if request.GetRouterConf() == nil {
		v.AddFieldViolation("router_info", "required field")
	} else {
		// host = 1
		if len(request.GetRouterConf().GetHost()) == 0 {
			v.AddFieldViolation("router_info.host", "required field")
		} else if matched, _ := regexp.MatchString(hostnamePattern, request.GetRouterConf().GetHost()); !matched {
			v.AddFieldViolation("router_info.host", "invalid hostname format")
		}
		// username = 2
		if len(request.GetRouterConf().GetUsername()) == 0 {
			v.AddFieldViolation("router_info.username", "required field")
		} else if len(request.GetRouterConf().GetUsername()) > maxUsernameLength {
			v.AddFieldViolation("router_info.username", "should be <= 255 characters")
		}
		// password = 3
		if len(request.GetRouterConf().GetPassword()) == 0 {
			v.AddFieldViolation("router_info.password", "required field")
		} else if len(request.GetRouterConf().GetPassword()) > maxPasswordLength {
			v.AddFieldViolation("router_info.password", "should be <= 255 characters")
		}

		r.routerConf = request.GetRouterConf()
	}

	return v.Err()
}
