package main

import (
	"fmt"
	"time"

	"github.com/Greateapot/roaure/internal/database"
)

func main() {
	db := database.NewDatabase("example.json")

	config := &database.Config{
		DownloadThreshold: 100 * database.KBit,
		PollInterval:      database.Time{Hours: 00, Minutes: 05},
		BadCountLimit:     3,
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
		return
	}

	// Output (example.json):
	// {
	// 	"downloadThreshold": 102400,
	// 	"pollInterval": {
	// 		"hours": 0,
	// 		"minutes": 5
	// 	},
	// 	"badCountLimit": 3,
	// 	"schedules": [
	// 		{
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

	if loaded, err := db.LoadConfig(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(loaded)
	}

	// Output (std):
	// &{12.50 KB {0 5} 3 [{New Schedule {10 0} {17 0} [Monday Wednesday Friday] true}]}
}
