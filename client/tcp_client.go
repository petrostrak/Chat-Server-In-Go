package client

import (
	"Chat-Server-In-Go/protocol"
	"fmt"
	"io"
	"log"
	"net"
)

type TcpChatServer struct {
	conn      net.Conn
	cmdReader *protocol.CommandReader
	cmdWriter *protocol.CommandeWriter
	name      string
	incoming  chan protocol.MessageCommand
}

func NewClient() *TcpChatServer {
	return &TcpChatServer{
		incoming: make(chan protocol.MessageCommand),
	}
}

func (c TcpChatServer) Dial(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		c.conn = conn
	}
	c.cmdReader = protocol.NewCommandReader(conn)
	c.cmdWriter = protocol.NewCommandWriter(conn)

	return err
}

func (c *TcpChatServer) Start() {
	for {
		cmd, err := c.cmdReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Read error %v", err)
		}

		if cmd != nil {
			switch v := cmd.(type) {
			case protocol.MessageCommand:
				c.incoming <- v
			default:
				log.Printf("Unknown command: %v", v)
			}
		}
	}
}

func (c *TcpChatServer) Close() {
	c.conn.Close()
}

func (c *TcpChatServer) Incoming() chan protocol.MessageCommand {
	return c.incoming
}

func (c *TcpChatServer) Send(command interface{}) error {
	return c.cmdWriter.Write(command)
}

func (c *TcpChatServer) SetName(name string) error {
	return c.Send(protocol.NameCommand{name})
}

func (c *TcpChatServer) SendMessage(message string) error {
	return c.Send(protocol.SendCommand{
		Message: message,
	})
}
