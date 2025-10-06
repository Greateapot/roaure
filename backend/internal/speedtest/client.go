package speedtest

import (
	"fmt"

	"github.com/BGrewell/go-iperf"
)

func RunClient(host string, port int) (float64, error) {
	c := iperf.NewClient(host)
	c.SetPort(port)
	c.SetReverse(true)
	c.SetTimeSec(10)
	c.SetInterval(2)

	if err := c.Start(); err != nil {
		return 0, fmt.Errorf("failed to start client: %v", err)
	}

	<-c.Done

	r := c.Report()
	if r.Error != "" {
		return 0, fmt.Errorf("%s", r.Error)
	}

	return r.End.SumReceived.BitsPerSecond, nil
}
