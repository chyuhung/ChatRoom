package main

import "ChatRoom/server"

func main() {
	var s *server.TcpServer
	s = server.NewServer()
	s.Listen(":27149")
	s.Start()
}
