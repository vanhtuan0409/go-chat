package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	chat "github.com/vanhtuan0409/go-chat"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go Chat Server"
	app.Usage = "Server for Chat application"
	app.Version = "v0.1.0"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: 8080,
			Usage: "Runing port of server",
		},
	}
	app.Action = handle
	app.Run(os.Args)
}

func handle(c *cli.Context) {
	port := c.Int("port")
	if port <= 0 {
		fmt.Println("require port value")
		return
	}

	server := chat.NewServer(port)
	sigTerm := make(chan os.Signal)

	go func() {
		log.Fatal(server.Listen())
	}()
	log.Printf("Server is running on :%d\n", port)

	signal.Notify(sigTerm, syscall.SIGTERM, syscall.SIGINT)
	<-sigTerm

	log.Printf("Server is shutting down\n")
	server.Shutdown()
}
