package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	// "github.com/TwiN/go-color"
	// "github.com/cqroot/prompt"
	// "github.com/cqroot/prompt/input"
	"github.com/eiannone/keyboard"

	"github.com/ylniss/go-net-game/messaging"
)

func main() {
	// net.Dial connects to a server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// diaplay menu and open keyboard read connection
	fmt.Println("[c] - connect to a new room\t[q] - quit")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	// wait for keyboard menu input
	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println("Error reading key. Please try again.")
			continue
		}

		if char == 'c' {
			err := sendMessageToServer(conn, &msg.ClientMessage{
				Event:     "connect_to_room",
				EventType: "menu",
				Data:      "",
			})
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

		} else if char == 'q' {
			os.Exit(0) // quit client with success
		}
	}

	// message, err = prompt.New().Ask(">").Input("")
	// if err != nil {
	// 	if errors.Is(err, prompt.ErrUserQuit) {
	// 		os.Exit(1)
	// 	}
	// 	return "", err
	// }
}

func sendMessageToServer(conn net.Conn, m *msg.ClientMessage) error {
	// 1) Marshal the struct to JSON
	data, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// 2) Write the JSON
	if _, err := conn.Write(data); err != nil {
		return fmt.Errorf("error writing to connection: %w", err)
	}

	return nil
}
