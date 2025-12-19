package main

import (
	"context"
	"time"

	"github.com/Greateapot/roaure/internal/database"
	"github.com/Greateapot/roaure/internal/led"
	"github.com/Greateapot/roaure/internal/monitor"
	"github.com/Greateapot/roaure/internal/router"
	"github.com/Greateapot/roaure/internal/speedtest"
)

func main() {
	rc := router.NewClient(
		&database.RouterConf{
			Host:     "192.168.1.1",
			Username: "admin",
			Password: "admin",
		},
		30*time.Second,
	)

	sc := speedtest.NewClient(
		&database.IperfServerConf{
			Host: "localhost",
			Port: 5201,
		},
	)

	led, _ := led.NewLED("/dev/gpiochip0", 7)

	m := monitor.NewMonitor(
		context.Background(),
		&database.MonitorConf{
			DownloadThreshold: 100 * database.KBit,
			PollInterval:      &database.Time{Hours: 00, Minutes: 05},
			BadCountLimit:     3,
			Schedules: []*database.Schedule{
				{
					Title:    "New Schedule",
					StartsAt: &database.Time{Hours: 10, Minutes: 00},
					EndsAt:   &database.Time{Hours: 17, Minutes: 00},
					Weekdays: []time.Weekday{
						time.Monday,
						time.Wednesday,
						time.Friday,
					},
					Enabled: true,
				},
			},
		},
		rc,
		sc,
		led,
	)

	m.Start()
	<-time.After(time.Hour)
	m.Stop()
}
