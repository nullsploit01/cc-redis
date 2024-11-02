package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
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
	reader := bufio.NewReader(conn)
	fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr().String())

	for {
		line, err := readRespCommand(reader)
		log.Println(line)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from client:", err)
			}
			break
		}

		response := processCommand(line)
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

func readRespCommand(reader *bufio.Reader) (string, error) {
	var fullCommand string
	arrayCount := -1
	readLines := 0

	for {
		part, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		fullCommand += part
		readLines++

		if arrayCount == -1 && strings.HasPrefix(part, "*") {
			count, err := strconv.Atoi(strings.TrimSpace(part[1:]))
			if err != nil {
				return "", err
			}
			arrayCount = count * 2 // each command has 2 parts, length and payload
		}

		if arrayCount != -1 && readLines >= arrayCount {
			break
		}
	}

	return strings.TrimRight(fullCommand, "\r\n"), nil
}
