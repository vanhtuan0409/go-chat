package chat

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
)

type Server struct {
	Port     int
	Context  context.Context
	cancelFn func()
	pool     *ConnPool
}

func NewServer(port int) *Server {
	ctx, cancelFn := context.WithCancel(context.Background())
	return &Server{
		Port:     port,
		Context:  ctx,
		cancelFn: cancelFn,
		pool:     NewPool(),
	}
}

func (s *Server) Listen() error {
	port := fmt.Sprintf(":%d", s.Port)
	l, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	// Event loop to accept request
	for {
		select {
		case <-s.Context.Done():
			return nil
		default:
			{
				conn, err := l.Accept()
				if err != nil {
					log.Printf("Cannot accept connection. ERR: %v", err)
					continue
				}
				go s.handle(conn)
			}
		}

	}
}

func (s *Server) Shutdown() {
	s.cancelFn()
}

func (s *Server) handle(tcpConn net.Conn) {
	log.Println("There is a new connection")
	tcpConn.Close()
	return

	// Parse connection
	conn, err := ParseConn(tcpConn)
	if err != nil {
		log.Printf("Cannot parse connection. ERR: %v\n", err)
		s.closeConn(conn, false)
		return
	}

	// register to connection pool
	if err := s.pool.Add(conn); err != nil {
		log.Printf("Cannot register to connection pool. ERR: %v\n", err)
		s.closeConn(conn, false)
		return
	}

	// event loop for received message
	for {
		msg, err := conn.ReadMessage()
		if err != nil {
			if err != io.EOF {
				log.Printf("Encounter an error while reading from connection. ERR: %v\n", err)
			}
			s.closeConn(conn, true)
			return
		}
		s.pool.Broadcast(msg)
	}
}

func (s *Server) closeConn(conn *Conn, shouldRemovePool bool) {
	conn.Close()
	if shouldRemovePool {
		s.pool.Remove(conn)
	}
	log.Println("A connection is closed")
}
