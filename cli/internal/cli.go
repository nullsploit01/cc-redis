package internal

import (
	"fmt"
	"net"
	"time"
)

type CLI struct {
	conn net.Conn
}

func NewCLI() *CLI {
	return &CLI{}
}

func (c *CLI) ConnectToServer(host, port string) error {
	addr := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		return fmt.Errorf("could not connect to server at %s", addr)
	}
	defer conn.Close()
	c.conn = conn
	return nil
}
