package main

import (
	"log"
	"net/http"

	server "github.com/besrabasant/pesto-clone-backend/server"
)

func main () {
	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	log.Println("Starting Server on port 8000")
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}