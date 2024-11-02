package internal

import (
	"fmt"
	"io"
	"net"
	"strings"
)

type Server struct {
	port string
}

func InitServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) StartServer() error {
	l, err := net.Listen("tcp", ":"+s.port)
	fmt.Printf("Listening on port %s\n", s.port)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr().String())

	for {
		packet := make([]byte, 4096)
		_, err := conn.Read(packet)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from client:", err)
			}
			break
		}

		fmt.Printf("Received packet from %s: %s\n", conn.RemoteAddr().String(), packet)
		response := processCommand(strings.TrimSpace(string(packet)))
		_, err = conn.Write([]byte(response + "\r\n"))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			break
		}
	}
}

func processCommand(command string) string {
	parts := strings.Fields(command)

	if len(parts) == 0 {
		return "-ERR no command received"
	}

	switch parts[0] {

	case "PING":
		return "+PONG"

	case "ECHO":
		if len(parts) < 2 {
			return "-ERR wrong number of arguments for 'ECHO' command"
		}
		return "+" + strings.Join(parts[1:], " ")

	case "SET":
		if len(parts) != 3 {
			return "-ERR wrong number of arguments for 'SET' command"
		}
		return "+OK"

	case "GET":
		if len(parts) != 2 {
			return "-ERR wrong number of arguments for 'GET' command"
		}
		return "$-1"

	default:
		return fmt.Sprintf("-ERR unknown command '%s'", parts[0])
	}
}
