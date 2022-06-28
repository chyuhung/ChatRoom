package server

type Server interface {
	Listen(address string) error
	Brodcast(command interface{}) error
	Start()
	Close()
}
