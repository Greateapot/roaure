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

type roaureServiceServer struct {
	db database.Database

	config          *database.Config
	routerClient    *router.Client
	speedtestClient *speedtest.Client
	monitor         *monitor.Monitor
}

func NewRoaureServiceServer(ctx context.Context, db database.Database) *roaureServiceServer {
	r := roaureServiceServer{db: db}

	if config, err := db.LoadConfig(); err != nil {
		grpclog.Fatal(err)
	} else {
		r.config = config
	}

	r.routerClient = router.NewClient(
		r.config.Router.Host,
		r.config.Router.Username,
		r.config.Router.Password,
		30*time.Second,
	)

	r.speedtestClient = speedtest.NewClient(
		r.config.Server.Host,
		r.config.Server.Port,
	)

	r.monitor = monitor.NewMonitor(
		ctx,
		r.config.DownloadThreshold,
		r.config.PollInterval,
		r.config.BadCountLimit,
		r.config.Schedules,
		r.routerClient,
		r.speedtestClient,
	)

	return &r
}

var _ roaurev1.RoaureServiceServer = (*roaureServiceServer)(nil)
