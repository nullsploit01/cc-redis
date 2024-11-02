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
		return fmt.Errorf("could not connect to ccredis-server at %s", addr)
	}
	defer conn.Close()
	c.conn = conn

	fmt.Printf("Connected to ccredis-server at %s\n", addr)
	for {
		c.Prompt()
		command := ""
		fmt.Scan(&command)
		err := c.SendCommand(command)
		if err != nil {
			return err
		}
	}
}

func (c *CLI) Prompt() {
	fmt.Print("> ")
}

func (c *CLI) SendCommand(command string) error {
	_, err := c.conn.Write([]byte(command))
	if err != nil {
		return err
	}
	return nil
}
