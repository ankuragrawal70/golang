package base

import "github.com/gorilla/websocket"

func PollResponse(conn *websocket.Conn){

	var reply string
	for{
		conn.ReadMessage()
	}
}
