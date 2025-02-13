package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (for development only)
	},
}

type Init_struct_2dGame struct {
	Type     string `json:"type"`
	Space_id int    `json:"space_id"`
	Username string `json:"username"`
	Height   int    `json:"height"`
	Width    int    `json:"Width"`
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

	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		err = json.Unmarshal(message, &Initilization)
		if err != nil {
			fmt.Println("Error Occered At a Scocket")
		}
		if Initilization.Type == INIT_SPACE {
			User, Room := UpdateUserRoomData(Initilization,conn)
			RoomsOnline[Room] = append(RoomsOnline[Room], User)
		}

		// Convert message to uppercase
		response := strings.ToUpper(string(message))

		// Send response back to client
		err = conn.WriteMessage(messageType, []byte(response))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func UpdateUserRoomData(Initilization Init_struct_2dGame ,conn *websocket.Conn) (PlayerObject,Room){
	var User PlayerObject
	User.Username = Initilization.Username
	User.Conn = conn
	var Room Room
	Room.roomId = Initilization.Space_id
	Room.Height = Initilization.Height
	Room.Width = Initilization.Height

    return User,Room
}



func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("WebSocket server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
