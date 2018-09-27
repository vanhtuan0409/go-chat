package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	chat "github.com/vanhtuan0409/go-chat"
)

const (
	ServerPort = 8080
)

func main() {
	server := chat.NewServer(ServerPort)
	sigTerm := make(chan os.Signal)

	go func() {
		log.Fatal(server.Listen())
	}()
	log.Printf("Server is running on :%d\n", ServerPort)

	signal.Notify(sigTerm, syscall.SIGTERM, syscall.SIGINT)
	<-sigTerm

	log.Printf("Server is shutting down\n")
	server.Shutdown()
}
