package server

import "net/http"

// AllRooms is the global hashmap  of all the Rooms
var AllRooms RoomMap

func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	response := []byte("Test CreateRoomRequestHandler")
	w.Write(response)
}

func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	response := []byte("Test JoinRoomRequestHandler")
	w.Write(response)
}