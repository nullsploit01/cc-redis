package internal

import (
	"bufio"
	"fmt"
	"io"
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
		command, err := readRespCommand(reader)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from client:", err)
			}
			break
		}

		if strings.TrimSpace(command) == "" {
			continue
		}
		response := processCommand(command)
		_, err = conn.Write([]byte(response + "\r\n"))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			break
		}
	}
}

func readRespCommand(reader *bufio.Reader) (string, error) {
	var result []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		line = strings.TrimRight(line, "\r\n")

		if strings.HasPrefix(line, "*") {
			numElements, err := strconv.Atoi(line[1:])
			if err != nil {
				return "", fmt.Errorf("invalid array count: %v", err)
			}
			elements, err := readArrayElements(reader, numElements)
			if err != nil {
				return "", err
			}
			result = append(result, elements...)
			break // Assuming array is always the entire command
		} else if strings.HasPrefix(line, "$") {
			bulkString, err := readBulkString(reader, line)
			if err != nil {
				return "", err
			}
			result = append(result, bulkString)
			continue
		} else {
			result = append(result, line)
			continue
		}
	}

	return strings.Join(result, " "), nil
}

func readArrayElements(reader *bufio.Reader, count int) ([]string, error) {
	var elements []string
	for i := 0; i < count; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimRight(line, "\r\n")

		if strings.HasPrefix(line, "$") {
			bulkString, err := readBulkString(reader, line)
			if err != nil {
				return nil, err
			}
			elements = append(elements, bulkString)
		} else {
			elements = append(elements, line)
		}
	}
	return elements, nil
}

func readBulkString(reader *bufio.Reader, initialLine string) (string, error) {
	size, err := strconv.Atoi(initialLine[1:])
	if err != nil {
		return "", fmt.Errorf("invalid bulk string size: %v", err)
	}

	if size < 0 {
		return "", nil // RESP null bulk string
	}

	valueLine, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimRight(valueLine, "\r\n"), nil
}

func processCommand(command string) string {
	parts := strings.Fields(command)

	if len(parts) == 0 {
		return "-ERR no command received"
	}

	switch strings.ToUpper(parts[0]) {

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
