package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/src"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (for development only)
	},
}

type Init_struct_2dGame struct {
	Type     string   `json:"type"`
	SpaceId  string   `json:"space_id"`
	Username string   `json:"username"`
	Height   string   `json:"height"`
	Width    string   `json:"width"`
	Movement Movement `json:"movement"`
}
type Movement struct {
	UserId  int `json:"userid"`
	SpaceId int `json:"spaceid"`
	X       int `json:"x"`
	Y       int `json:"y"`
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")
	var Initilization Init_struct_2dGame
	var Room src.Room
	var User src.PlayerObject

	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		err = json.Unmarshal(message, &Initilization)
		fmt.Println(Initilization)

		if err != nil {
			fmt.Println("Error Occered At a Scocket")
			return
		}
		if Initilization.Type == src.INIT_SPACE {
			User, Room = UpdateUserRoomData(Initilization, conn)

			src.RoomsOnline[Room] = append(src.RoomsOnline[Room], User)

		} else if Initilization.Type == src.MOVE {
			AllRoomPlayers := src.RoomsOnline[Room]
			for room := 0; room < len(AllRoomPlayers); room++ {
				singleClient := AllRoomPlayers[room]
				responce, _ := json.Marshal(Initilization.Movement)
				singleClient.Conn.WriteMessage(messageType, []byte(responce))
			}
			response := strings.ToUpper("DATA UPDATED TO ALL USERS")

			// Send response back to client
			err = conn.WriteMessage(messageType, []byte(response))
			if err != nil {
				log.Println("Write error:", err)
				break
			}

		}
	}
}

func UpdateUserRoomData(init Init_struct_2dGame, conn *websocket.Conn) (src.PlayerObject, src.Room) {
	var User src.PlayerObject
	User.Username = init.Username
	User.Conn = conn

	var Room src.Room
	Room.RoomId = init.SpaceId

	// Convert Height and Width from string to int
	height, err := strconv.Atoi(init.Height)
	if err != nil {
		log.Println("Error converting height:", err)
		return User, Room // Return empty structs in case of error
	}
	Room.Height = height

	width, err := strconv.Atoi(init.Width)
	if err != nil {
		log.Println("Error converting width:", err)
		return User, Room // Return empty structs in case of error
	}
	Room.Width = width

	return User, Room
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("WebSocket server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
