package service

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	Msgch      chan Message
}

type Message struct {
	Conn    net.Conn
	Payload []byte
}

func NewTcpServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		Msgch:      make(chan Message, 20),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch
	close(s.Msgch)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		log.Println("New connection to the server:", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()

	for {
		// Put buf inside loop to prevent being reused by other process
		buf := make([]byte, 2048)

		n, err := conn.Read(buf)
		if err != nil {
			log.Println("read error:", err)
			return
		}

		s.Msgch <- Message{
			Conn:    conn,
			Payload: buf[:n],
		}
	}
}
