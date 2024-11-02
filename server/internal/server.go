package internal

import (
	"fmt"
	"io"
	"net"
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
				continue
			}
			break
		}

		fmt.Printf("Received packet from %s: %s\n", conn.RemoteAddr().String(), packet)
	}

}
