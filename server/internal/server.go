package internal

import (
	"fmt"
	"io"
	"net"
)

type Server struct {
	port string
}

func InitServer(port string) {
	s := &Server{port: port}
	s.StartServer()
}

func (s *Server) StartServer() {
	l, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr().String())

	packet := make([]byte, 4096)
	tmp := make([]byte, 4096)

	for {
		_, err := conn.Read(packet)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from client:", err)
				continue
			}
			break
		}
		packet = append(packet, tmp...)
	}
}
