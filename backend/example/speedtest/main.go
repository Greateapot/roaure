package main

import (
	"fmt"

	"github.com/Greateapot/roaure/internal/database"
	"github.com/Greateapot/roaure/internal/speedtest"
)

func main() {
	if bps, err := speedtest.RunClient("localhost", 5201); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Download speed: %s", database.DataSize(bps).String())
	}
}
