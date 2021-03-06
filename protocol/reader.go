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
func (r *Reader) Read() (interface{}, error) {
	cmd, err := r.reader.ReadString(' ')
	if err != nil {
		return nil, err
	}
	switch cmd {
	//ReadString reads until the first occurrence of delim in the input,
	//returning a string containing the data up to and including the delimiter.
	case "SEND ":
		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return SendCmd{message[:len(message)-1]}, nil
	case "NAME ":
		name, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return NameCmd{name[:len(name)-1]}, nil
	case "MESS ":
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
		log.Printf("UnknownCommand:%v", cmd)
	}
	return nil, UnknownCommand
}

func (r *Reader) ReadAll() ([]interface{}, error) {
	var commands []interface{}
	for {
		command, err := r.Read()
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
