package main

import (
	"fmt"
	"net"
)

func main() {
	// net.Dial connects to a server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	message := "Hello, Server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}
	fmt.Println("Sent message to server:", message)
}
