package main

import (
	"fmt"
	"time"

	"github.com/Greateapot/roaure/internal/router"
)

func main() {
	c := router.NewClient("192.168.1.1", "admin", "admin", 5*time.Minute)
	if err := c.Reboot(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("done")
	}
}
