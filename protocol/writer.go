package protocol

import (
	"fmt"
	"io"
)

//实现io.Writer接口
type Writer struct {
	writer io.Writer
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{writer: writer}
}

func (w *Writer) WriteString(msg string) error {
	_, err := w.writer.Write([]byte(msg))
	return err
}

//实现Write方法
func (w *Writer) Write(command interface{}) error {
	var err error
	switch v := command.(type) {
	case SendCmd:
		err = w.WriteString(fmt.Sprintf("SEND %v\n", v.Message))
	case NameCmd:
		err = w.WriteString(fmt.Sprintf("NAME %v\n", v.Name))
	case MessCmd:
		err = w.WriteString(fmt.Sprintf("MESS %v %v\n", v.Name, v.Message))
	default:
		err = UnkownCommand
	}
	return err
}
