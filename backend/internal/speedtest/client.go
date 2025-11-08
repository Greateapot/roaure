package speedtest

import (
	"fmt"

	"github.com/BGrewell/go-iperf"
	"github.com/Greateapot/roaure/internal/database"
)

type Client struct {
	IperfServerConf *database.IperfServerConf

	iperfClient *iperf.Client
}

func NewClient(iperfServerConf *database.IperfServerConf) *Client {
	c := Client{IperfServerConf: iperfServerConf}

	c.SetupClient()

	return &c
}

func (c *Client) SetupClient() {
	c.Stop()

	c.iperfClient = iperf.NewClient(c.IperfServerConf.Host)
	c.iperfClient.SetPort(int(c.IperfServerConf.Port))
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
