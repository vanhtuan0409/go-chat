package chat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Conn struct {
	Name   string
	conn   net.Conn
	reader *bufio.Reader
}

type Message struct {
	Name    string
	Message string
}

func (m *Message) String() string {
	return fmt.Sprintf("%s: %s", m.Name, m.Message)
}

func ParseConn(conn net.Conn) (*Conn, error) {
	reader := bufio.NewReader(conn)

	// reader name from connection
	name, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return &Conn{
		Name:   strings.Trim(name, "\n\t "),
		conn:   conn,
		reader: reader,
	}, nil
}

func (c *Conn) ReadMessage() (*Message, error) {
	msg, err := c.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	return &Message{
		Name:    c.Name,
		Message: msg,
	}, nil
}

func (c *Conn) WriteMessage(m *Message) error {
	_, err := c.conn.Write([]byte(m.String()))
	return err
}

func (c *Conn) Close() error {
	return c.conn.Close()
}
