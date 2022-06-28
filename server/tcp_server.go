package server

import (
	"ChatRoom/protocol"
	"net"
)

type client struct {
	conn   net.Conn
	name   string
	writer *protocol.Writer
}
