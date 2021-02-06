package main

import "Chat-Server-In-Go/server"

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen(":3333")

	// start the server
	s.Start()
}
