package chat

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type Client struct {
	Name       string
	ServerIP   string
	ServerPort int
}

func NewClient(name string, serverIP string, serverPort int) *Client {
	return &Client{
		Name:       name,
		ServerIP:   serverIP,
		ServerPort: serverPort,
	}
}

func (c *Client) Connect() error {
	addr := fmt.Sprintf("%s:%d", c.ServerIP, c.ServerPort)
	log.Printf("Trying to connect to %s as %s\n", addr, c.Name)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.Write([]byte(fmt.Sprintf("%s\n", c.Name))); err != nil {
		return err
	}

	go io.Copy(conn, os.Stdout)
	io.Copy(os.Stdin, conn)
	return nil
}
