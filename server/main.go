package main

import (
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Player struct {
	Id     uuid.UUID
	RoomId uuid.UUID
}

type Room struct {
	Id      uuid.UUID
	Player1 *Player
	Player2 *Player
}

var (
	players []*Player
	rooms   []*Room
)

func main() {
	// net.Listen creates a server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")
	fmt.Println()

	connectedPlayers := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		connectedPlayers++

		// every second player connecting join existing room
		if connectedPlayers%2 == 0 {
			awaitingRoom, found := lo.Find(rooms, func(r *Room) bool {
				return r.Player1 != nil && r.Player2 == nil
			})
			if found {
				go handleConnection(conn, awaitingRoom)
			} else {
				fmt.Println("Couldn't find free room")
			}
		} else {
			go handleConnection(conn, nil)
		}
	}
}

func handleConnection(conn net.Conn, room *Room) {
	defer conn.Close()

	connectedPlayerId := uuid.New()
	if room == nil {
		newRoom := createNewRoom()
		newPlayer := &Player{Id: connectedPlayerId, RoomId: newRoom.Id}
		players = append(players, newPlayer)
		newRoom.Player1 = newPlayer
		fmt.Printf("Player [%s] created the room [%s]\n", newPlayer.Id, newRoom.Id)
	} else {
		// room already exists
		newPlayer := &Player{Id: connectedPlayerId, RoomId: room.Id}
		players = append(players, newPlayer)
		room.Player2 = newPlayer
		fmt.Printf("Player [%s] joined the room [%s]\n", newPlayer.Id, room.Id)
	}

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	fmt.Println("Received data:", string(buffer))
}

func createNewRoom() *Room {
	newRoom := &Room{Id: uuid.New()}
	rooms = append(rooms, newRoom)
	return newRoom
}
