package main

import (
	"fmt"
	"log"
	"os"

	chat "github.com/vanhtuan0409/go-chat"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go Chat Client"
	app.Usage = "Client for Chat application"
	app.Version = "v0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "",
			Usage: "User name when chat. Should be unique",
		},
		cli.StringFlag{
			Name:  "address, a",
			Value: "localhost",
			Usage: "Address name of server",
		},
		cli.IntFlag{
			Name:  "port, p",
			Value: 8080,
			Usage: "Runing port of server",
		},
	}
	app.Action = func(c *cli.Context) {
		name := c.String("name")
		if name == "" {
			fmt.Println("require name value")
			return
		}

		address := c.String("address")
		if address == "" {
			fmt.Println("require address value")
			return
		}

		port := c.Int("port")
		if port <= 0 {
			fmt.Println("require port value")
			return
		}

		client := chat.NewClient(name, address, port)
		log.Fatal(client.Connect())
	}

	app.Run(os.Args)
}
