package server

// ChatServer interface
type ChatServer interface {
	Listen(address string) error
	Broadcast(command interface{}) error
	Start()
	Close()
}
