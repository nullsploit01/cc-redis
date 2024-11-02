package internal

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type CLI struct {
	conn   net.Conn
	reader *bufio.Reader
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
	c.reader = bufio.NewReader(conn)

	fmt.Printf("Connected to ccredis-server at %s\n", addr)
	c.Start()
	return nil
}

func (c *CLI) Start() error {
	defer c.conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return err
			}
			break
		}

		command := scanner.Text()
		if strings.ToUpper(command) == "QUIT" {
			fmt.Println("Bye!")
			break
		}

		if err := c.SendCommand(command); err != nil {
			return fmt.Errorf("could not send command: %s", err)
		}

		response, err := c.ReadResponse()
		if err != nil {
			return fmt.Errorf("could not read response: %s", err)
		}

		fmt.Println(response)
	}

	return nil
}

func (c *CLI) SendCommand(command string) error {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return nil
	}

	respCommand := fmt.Sprintf("*%d\r\n", len(parts))
	for _, part := range parts {
		respCommand += fmt.Sprintf("$%d\r\n%s\r\n", len(part), part)
	}

	_, err := c.conn.Write([]byte(respCommand))
	return err
}

func (c *CLI) ReadResponse() (string, error) {
	resp, err := c.reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading response: %s", err)
	}

	resp = strings.TrimSuffix(resp, "\r\n")
	if resp == "" {
		return "", fmt.Errorf("empty response from server")
	}

	switch resp[0] {
	case '+':
		return strings.TrimPrefix(resp, "+"), nil

	case '-':
		return "Error: " + strings.TrimPrefix(resp, "-"), nil

	case ':':
		return strings.TrimPrefix(resp, ":"), nil

	case '$':
		if len(resp) > 1 && resp[1] == '-' {
			return "Nil", nil
		}

		if content, err := c.reader.ReadString('\n'); err == nil {
			return strings.TrimSuffix(content, "\r\n"), nil
		}

	case '*':
		return strings.TrimPrefix(resp, "*"), nil
	}

	return "Unrecognized response: " + resp, nil
}
