package main

import (
	"flag"
	"log"

	"ChatRoom/client"
	"ChatRoom/gui"
)

func main() {
	address := flag.String("server", "127.0.0.1:27149", "address of server")
	flag.Parse()
	client := client.NewClient()
	err := client.Dial(*address)

	if err != nil {
		log.Fatal("Error when connect to server", err)
	}

	defer client.Close()

	go client.Start()
	gui.StartUI(client)
}
