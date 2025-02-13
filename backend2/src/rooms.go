package main

import 	"github.com/gorilla/websocket"
// {
// 	"type":"init_space",
// 	"space_id":"1",
// 	"username":"Balaiah"
//   }

type Room struct{
	roomId int
	Width int
	Height int
}

type Movement struct{
	UserId int `json:"userid"`
	X int `json:"x"`
	Y int `json:"y"`
}

type PlayerObject struct{
	Conn *websocket.Conn
	Username string
}


var RoomsOnline map[Room][]PlayerObject