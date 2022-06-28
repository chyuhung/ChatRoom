package protocol

import (
	"bufio"
	"io"
	"log"
)

//实现io.Reader
type Reader struct {
	reader *bufio.Reader
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{reader: bufio.NewReader(reader)}
}
func (r *Reader) Reade() (interface{}, error) {
	cmd, err := r.reader.ReadString(' ')
	if err != nil {
		return nil, err
	}
	switch cmd {
	case "SEND":
		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return SendCmd{message[:len(message)-1]}, nil
	case "NAME":
		name, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return NameCmd{name[:len(name)-1]}, nil
	case "MESS":
		user, err := r.reader.ReadString(' ')
		if err != nil {
			return nil, err
		}
		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return MessCmd{user[:len(user)-1], message[:len(message)-1]}, nil
	default:
		log.Printf("UnknowCommand:v%", cmd)
	}
	return nil, UnkownCommand
}

func (r *Reader) ReadAll() ([]interface{}, error) {
	commands := []interface{}{}
	for {
		command, err := r.Reade()
		if command != nil {
			commands = append(commands, command)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return commands, err
		}
	}
	return commands, nil
}