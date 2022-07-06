package server

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Participant describes a Participant in the Room
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// RoomMap describes the Room with all the participants in a Room
type RoomMap struct {
	Mutex sync.Mutex
	Map map[string][]Participant
}

// Initialize the RoomMap
func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

// Return all the Participants of the Room
func (r *RoomMap) Get(roomID string) []Participant {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.Map[roomID]
}

// CreateRoom will create a New Room and return the Rooom ID
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	roomID := uuid.New().String()

	r.Map[roomID] = []Participant{}

	return roomID
}

// InsertToRoom will create a new particpant and insert it to the Room Map
func(r *RoomMap) InsertToRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p:= Participant{host, conn}
	r.Map[roomID] = append(r.Map[roomID], p)
}

// DeleteRooom will delete the Room with the provided RoomdID
func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)
}
