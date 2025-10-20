package main

import (
	"context"
	"time"

	"github.com/Greateapot/roaure/internal/database"
	"github.com/Greateapot/roaure/internal/monitor"
	"github.com/Greateapot/roaure/internal/router"
	"github.com/Greateapot/roaure/internal/speedtest"
)

func main() {
	config := &database.Config{
		DownloadThreshold: 100 * database.KBit,
		PollInterval:      database.Time{Hours: 00, Minutes: 05},
		BadCountLimit:     3,
		Server: database.Server{
			Host: "localhost",
			Port: 5201,
		},
		Router: database.Router{
			Host:     "192.168.1.1",
			Username: "admin",
			Password: "admin",
		},
		Schedules: []database.Schedule{
			{
				Title:    "New Schedule",
				StartsAt: database.Time{Hours: 10, Minutes: 00},
				EndsAt:   database.Time{Hours: 17, Minutes: 00},
				Weekdays: []time.Weekday{
					time.Monday,
					time.Wednesday,
					time.Friday,
				},
				Enabled: true,
			},
		},
	}

	rc := router.NewClient(
		config.Router.Host,
		config.Router.Username,
		config.Router.Password,
		30*time.Second,
	)

	sc := speedtest.NewClient(
		config.Server.Host,
		config.Server.Port,
	)

	m := monitor.NewMonitor(
		context.Background(),
		config.DownloadThreshold,
		config.PollInterval,
		config.BadCountLimit,
		config.Schedules,
		rc,
		sc,
	)

	m.Start()
	<-time.After(time.Hour)
	m.Stop()
}
