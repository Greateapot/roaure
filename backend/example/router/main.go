package main

import (
	"fmt"
	"time"

	"github.com/Greateapot/roaure/internal/database"
	"github.com/Greateapot/roaure/internal/router"
)

func main() {
	c := router.NewClient(&database.RouterConf{
		Host:     "192.168.1.1",
		Username: "admin",
		Password: "admin",
	}, 30*time.Second)
	if err := c.Reboot(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("done")
	}
}
