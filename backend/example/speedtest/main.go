package main

import (
	"fmt"

	"github.com/Greateapot/roaure/internal/database"
	"github.com/Greateapot/roaure/internal/speedtest"
)

func main() {
	c := speedtest.NewClient(&database.IperfServerConf{
		Host: "localhost",
		Port: 5201,
	})

	if bps, err := c.Start(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Download speed: %s", database.DataSize(bps).String())
	}
}
