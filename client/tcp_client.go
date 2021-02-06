package client

import (
	"Chat-Server-In-Go/protocol"
	"net"
)

type TcpChatServer struct {
	conn      net.Conn
	cmdReader *protocol.CommandReader
	cmdWriter *protocol.CommandeWriter
	name      string
	incoming  chan protocol.MessageCommand
}
