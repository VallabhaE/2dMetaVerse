package src

import 	"github.com/gorilla/websocket"

// {
// 	"type":"init_space",
// 	"space_id":"1",
// 	"username":"Balaiah",
// 	"height":"24",
// 	"Width":"24"
//   }
type Room struct{
	RoomId string
	Width int
	Height int
}


// Sample Json
// {
// 	"userId":"1",
// 	"spaceid":"1"
// 	"x":"1",
// 	"y":"2"
//   }

type PlayerObject struct{
	Conn *websocket.Conn
	Username string
}


var RoomsOnline =  make(map[Room][]PlayerObject)