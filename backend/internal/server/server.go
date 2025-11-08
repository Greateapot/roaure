package server

import (
	"context"
	"time"

	"github.com/Greateapot/roaure/internal/database"
	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"github.com/Greateapot/roaure/internal/monitor"
	"github.com/Greateapot/roaure/internal/router"
	"github.com/Greateapot/roaure/internal/speedtest"
	"google.golang.org/grpc/grpclog"
)

const (
	defaultTimeout = 30 * time.Second
)

type roaureServiceServer struct {
	db database.Database

	config  *database.RoaureConf
	monitor *monitor.Monitor
}

func NewRoaureServiceServer(ctx context.Context, db database.Database) *roaureServiceServer {
	r := roaureServiceServer{db: db}

	if config, err := r.db.LoadConfig(); err == nil {
		r.config = config
	} else if config, err := r.db.NewConfig(); err == nil {
		r.config = config
	} else {
		grpclog.Fatalln(err)
	}

	r.monitor = monitor.NewMonitor(
		ctx,
		r.config.MonitorConf,
		router.NewClient(
			r.config.RouterConf,
			defaultTimeout,
		), speedtest.NewClient(
			r.config.IperfServerConf,
		),
	)

	return &r
}

var _ roaurev1.RoaureServiceServer = (*roaureServiceServer)(nil)
