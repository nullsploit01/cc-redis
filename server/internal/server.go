package internal

import "net"

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
}
