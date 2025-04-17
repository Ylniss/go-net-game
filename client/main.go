package main

import (
	"fmt"
	"net"
	"os"

	"github.com/TwiN/go-color"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
	"github.com/eiannone/keyboard"

	"github.com/ylniss/go-net-game/msg"
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
	fmt.Printf("[c] - connect to a new room%t[q] - quit\n\n")

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
            err := sendMessageToServer(&msg.ClientMessage{Event: "connect_to_room", EventType: "menu", Data: ""})
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

func sendMessageToServer(msg *msg.ClientMessage) : err {
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return err
	}

    return _
	// fmt.Println("Sent message to server:", message)
    
}
