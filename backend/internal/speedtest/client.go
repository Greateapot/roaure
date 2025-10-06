package speedtest

import (
	"fmt"

	"github.com/BGrewell/go-iperf"
)

type Client struct {
	Host string
	Port int

	iperfClient *iperf.Client
}

func NewClient(host string, port int) *Client {
	c := Client{
		Host: host,
		Port: port,
	}

	c.SetupClient()

	return &c
}

func (c *Client) SetupClient() {
	c.Stop()

	c.iperfClient = iperf.NewClient(c.Host)
	c.iperfClient.SetPort(c.Port)
	c.iperfClient.SetReverse(true)
	c.iperfClient.SetTimeSec(10)
	c.iperfClient.SetInterval(2)
}

func (c *Client) Start() (float64, error) {
	if err := c.iperfClient.Start(); err != nil {
		return 0, fmt.Errorf("failed to start client: %v", err)
	}

	<-c.iperfClient.Done

	r := c.iperfClient.Report()
	if r.Error != "" {
		return 0, fmt.Errorf("%s", r.Error)
	}

	return r.End.SumReceived.BitsPerSecond, nil
}

func (c *Client) Stop() {
	if c.iperfClient != nil {
		c.iperfClient.Stop()
	}
}
