package main

import (
	"Chat-Server-In-Go/client"
	"Chat-Server-In-Go/tui"
	"flag"
	"log"
)

func main() {
	address := flag.String("server", "", "Which server to connect to")
	flag.Parse()

	client := client.NewClient()
	err := client.Dial(*address)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// start the client to listen for incoming message
	go client.Start()

	tui.StartUi(client)
}
