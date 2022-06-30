package protocol

import "errors"

var UnknownCommand = errors.New("UnknownCommand")

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
