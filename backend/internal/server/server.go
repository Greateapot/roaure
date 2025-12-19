package server

import (
	"context"
	"time"

	"github.com/Greateapot/roaure/internal/database"
	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"github.com/Greateapot/roaure/internal/led"
	"github.com/Greateapot/roaure/internal/monitor"
	"github.com/Greateapot/roaure/internal/router"
	"github.com/Greateapot/roaure/internal/speedtest"
	"google.golang.org/grpc/grpclog"
)

const (
	defaultTimeout = 30 * time.Second
)

type roaureServiceServer struct {
	Database database.Database

	config  *database.RoaureConf
	monitor *monitor.Monitor
	led     *led.LED
}

func NewRoaureServiceServer(
	ctx context.Context,
	database database.Database,
	ledChip string,
	ledlLineOffset int,
) *roaureServiceServer {
	r := roaureServiceServer{Database: database}

	if config, err := database.LoadConfig(); err == nil {
		r.config = config
	} else if config, err := database.NewConfig(); err == nil {
		r.config = config
	} else {
		grpclog.Fatalln(err)
	}

	led, err := led.NewLED(ledChip, ledlLineOffset)
	if err != nil {
		grpclog.Fatalln(err)
	}

	r.monitor = monitor.NewMonitor(
		ctx,
		r.config.MonitorConf,
		router.NewClient(
			r.config.RouterConf,
			defaultTimeout,
		),
		speedtest.NewClient(
			r.config.IperfServerConf,
		),
		led,
	)

	return &r
}

var _ roaurev1.RoaureServiceServer = &roaureServiceServer{}
