package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Greateapot/roaure/internal/database"
)

func main() {
	db := database.NewDatabase("example.json")

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

	if err := db.DumpConfig(config); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// Output (example.json):
	// {
	// 	"downloadThreshold": 102400,
	// 	"pollInterval": {
	// 		"hours": 0,
	// 		"minutes": 5
	// 	},
	// 	"badCountLimit": 3,
	// 	"server": {
	// 		"host": "localhost",
	// 		"port": 5201
	// 	},
	// 	"router": {
	// 		"host": "192.168.1.1",
	// 		"username": "admin",
	// 		"password": "admin"
	// 	},
	// 	"schedules": [
	// 		{
	// 			"ID": "00000000-0000-0000-0000-000000000000",
	// 			"title": "New Schedule",
	// 			"starts_at": {
	// 				"hours": 10,
	// 				"minutes": 0
	// 			},
	// 			"ends_at": {
	// 				"hours": 17,
	// 				"minutes": 0
	// 			},
	// 			"weekdays": [
	// 				1,
	// 				3,
	// 				5
	// 			],
	// 			"enabled": true
	// 		}
	// 	]
	// }

	loaded, err := db.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(loaded)
}
