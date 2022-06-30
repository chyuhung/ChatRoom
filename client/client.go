package client

import "ChatRoom/protocol"

//客户端的主要功能如下：
//▪连接服务器。
//▪使用用户名登录。
//▪发送消息。
//▪接收其他人发送的消息。

type Client interface {
	Dial(address string) error
	Start()
	Close()
	Send(command interface{}) error
	SendMess(message string) error
	SetName(name string) error
	InComing() chan protocol.MessCmd
}
