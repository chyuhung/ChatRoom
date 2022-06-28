package protocol

import "errors"

var UnkownCommand = errors.New("Unkown Command")

type SendCmd struct {
	Message string
}
type NameCmd struct {
	Name string
}
type MessCmd struct {
	Name    string
	Message string
}
