package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	dictClient()
}

const port = ":4546"

func dictClient() {
	address := "127.0.0.1" + port
	connection, connectionError := net.Dial("tcp", address)
	if connectionError != nil {
		fmt.Errorf("Connection err: %s\n", connectionError)
		return
	}
	defer connection.Close()
	for {
		var promoKey string
		fmt.Println("Enter the keyword to get the promo code:")
		if _, scanError := fmt.Scanln(&promoKey); scanError != nil {
			fmt.Errorf("Scan err: %s\n", scanError)
			continue
		}
		bytes := []byte(promoKey)
		if _, writeError := connection.Write(bytes); writeError != nil {
			fmt.Errorf("Write err: %s\n", writeError)
			return
		}

		fmt.Println("Promo code response:")
		connection.SetReadDeadline(time.Now().Add(time.Second * 5))
		for {
			buffer := make([]byte, 1024)
			bufferSize, getPromoErr := connection.Read(buffer)
			if getPromoErr != nil {
				fmt.Errorf("Read err: %s\n", getPromoErr)
				break
			}
			promo := string(buffer[:bufferSize])
			fmt.Print(promo, "\n")
			connection.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
		}
	}
}
