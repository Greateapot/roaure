package router

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Greateapot/roaure/internal/database"
)

const (
	sessionKeyPattern = `.*var\s*sessionKey='([0-9]+)'.*`
	rebootMessage     = `The Broadband Router is rebooting`
)

// Клиент роутера
type Client struct {
	RouterConf *database.RouterConf

	client http.Client
}

func NewClient(
	routerConf *database.RouterConf,
	timeout time.Duration,
) *Client {
	c := Client{RouterConf: routerConf}

	c.client = http.Client{Timeout: timeout}

	return &c
}

func (c *Client) obtainSessionKey() (string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://%s/resetrouter.html", c.RouterConf.Host),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("create req: %v", err)
	}

	req.SetBasicAuth(c.RouterConf.Username, c.RouterConf.Password)
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("make req: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code is not OK: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("invalid body: %v", err)
	}

	matches := regexp.MustCompile(sessionKeyPattern).FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "", fmt.Errorf("unnable to match sessionKey")
	}

	return matches[1], nil
}

func (c *Client) reboot(sessionKey string) error {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://%s/rebootinfo.cgi?sessionKey=%s", c.RouterConf.Host, sessionKey),
		nil,
	)
	if err != nil {
		return fmt.Errorf("create req: %v", err)
	}

	req.SetBasicAuth(c.RouterConf.Username, c.RouterConf.Password)
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("make req: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is not OK: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("invalid body: %v", err)
	}

	if !strings.Contains(string(body), rebootMessage) {
		return fmt.Errorf("unnable to reboot")
	}

	return nil
}

func (c *Client) Reboot() error {
	if sessionKey, err := c.obtainSessionKey(); err != nil {
		return err
	} else {
		return c.reboot(sessionKey)
	}
}
