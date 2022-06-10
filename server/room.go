package server

import (
	"sync"
	//"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

// Participant describes a single entity in the hashmap
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// RoomMap is the main hashmap [roomID string] -> [[]Participant]
type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

// Initialises the RoomMap struct
func (r *RoomMap) Init() {

	r.Map = make(map[string][]Participant)
}

// will return  the array of participants in the room
func (r *RoomMap) GetParticipantsInRoom(roomID string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[roomID]
}

// generate a unique ID ande return it -> insert it in the hashmap
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)
	r.Map[roomID] = []Participant{}

	return roomID
}

// will create a participant and add it in the hashmap
func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}

	log.Println("Inserting into Room with RoomID: ", roomID)
	r.Map[roomID] = append(r.Map[roomID], p)
}

// Deletes the room with the roomID
func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)
	log.Print("Im in ther and delete all. Now r.MAP this: ", r.Map)
}


