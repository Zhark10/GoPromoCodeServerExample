package handler

import (
	"fmt"
	"net"
)

func NewConnectionHandle(connection net.Conn) {
	defer connection.Close()

	for {
		data := make([]byte, 1024*2)
		size, readError := connection.Read(data)

		if size == 0 || readError != nil {
			fmt.Errorf("Read err: %s\n", readError)
			break
		}

		productIdToSearch := data[:size]
		target, ok := PromoDict[string(productIdToSearch)]
		message := "unknown"
		if ok {
			message = target
		}
		connection.Write([]byte(message))
	}
}
