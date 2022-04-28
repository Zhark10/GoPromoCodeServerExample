package main

import (
	"fmt"
	"net"
)

type productId = string

const ListenPort = ":4546"

func main() {
	startServer()
}

func startServer() {
	listener, error := net.Listen("tcp", ListenPort)
	if error != nil {
		fmt.Errorf("Listen err: %s\n", error)
		return
	}
	fmt.Printf("Listen on port%s\n", ListenPort)
	defer listener.Close()
	for {
		newConnection, connectionError := listener.Accept()
		if connectionError != nil {
			fmt.Errorf("Connection err: %s\n", error)
			newConnection.Close()
			continue
		}
		go NewConnectionHandle(newConnection)
	}
}
