package client

import (
	"ChatRoom/protocol"
	"io"
	"log"
	"net"
	"time"
)

type TcpClient struct {
	conn      net.Conn
	cmdReader *protocol.Reader
	cmdWriter *protocol.Writer
	name      string
	inComing  chan protocol.MessCmd
}

func NewClient() *TcpClient {
	return &TcpClient{
		inComing: make(chan protocol.MessCmd),
	}
}
func (c *TcpClient) Dial(address string) error {
	log.Println(address)
	conn, err := net.Dial("tcp", address)
	if err == nil {
		c.conn = conn
	} else {
		log.Println("dial error!")
		return err
	}
	c.cmdReader = protocol.NewReader(conn)
	c.cmdWriter = protocol.NewWriter(conn)
	return err
}
func (c *TcpClient) Start() {
	log.Println("starting client.")
	time.Sleep(4 * time.Second)
	for {
		cmd, err := c.cmdReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Read error %v", err)
		}
		if cmd != nil {
			switch v := cmd.(type) {
			case protocol.MessCmd:
				c.inComing <- v
			default:
				log.Printf("Unknown command %v", v)
			}
		}
	}
}
func (c *TcpClient) Close() {
	c.conn.Close()

}
func (c *TcpClient) Send(command interface{}) error {
	return c.cmdWriter.Write(command)
}
func (c *TcpClient) SendMess(message string) error {
	return c.Send(protocol.SendCmd{Message: message})
}
func (c *TcpClient) SetName(name string) error {
	return c.Send(protocol.NameCmd{Name: name})
}
func (c *TcpClient) InComing() chan protocol.MessCmd {
	return c.inComing
}
