package main

import (
	"log"

	chat "github.com/vanhtuan0409/go-chat"
)

func main() {
	client := chat.NewClient("tuanva", "localhost", 8080)
	log.Fatal(client.Connect())
}
