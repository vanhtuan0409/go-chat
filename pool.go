package chat

import (
	"errors"
	"fmt"
	"log"
)

type ConnPool struct {
	connections map[string]*Conn
}

func NewPool() *ConnPool {
	return &ConnPool{
		connections: map[string]*Conn{},
	}
}

func (p *ConnPool) Add(conn *Conn) error {
	if _, ok := p.connections[conn.Name]; ok {
		msg := fmt.Sprintf("Connection with name %s already exist in pool", conn.Name)
		return errors.New(msg)
	}
	p.connections[conn.Name] = conn
	log.Printf("Connection with name %s registered to pool\n", conn.Name)
	return nil
}

func (p *ConnPool) Remove(conn *Conn) {
	delete(p.connections, conn.Name)
	log.Printf("Connection with name %s deregistered from pool\n", conn.Name)
}

func (p *ConnPool) Broadcast(m *Message) {
	log.Println("Broadcast message to all connection")
	for _, conn := range p.connections {
		conn.WriteMessage(m)
	}
}
