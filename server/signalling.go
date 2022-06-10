package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v4"
)

// is the global hashmap for the server
var AllRooms RoomMap
var URLDatabase = "postgres://postgres:pass@localhost:5432/Fart"

type js_LP struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//Check of correct enter Login and Password
func LoginReguestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	decoder := json.NewDecoder(r.Body)
	var user js_LP 
	err := decoder.Decode(&user)
	if err != nil {
		log.Print("JSON faild: ", err)
	}

	//log.Print(user.Login)

	type resp struct {
		ID uint16 `json:"id"`
		Login string`json:"login"`
		Access bool `json:"Access"`
	}

	connPG, err := pgx.Connect(context.Background(), URLDatabase)

	if err != nil {
		log.Print("Unable to connect to database:", err)
	}
	defer connPG.Close(context.Background())

	var id_user uint16
	err = connPG.QueryRow(
		context.Background(),
		"select id from logins where login=$1 and pass=$2 ", user.Login, user.Password).Scan(&id_user)
	if err != nil {
		log.Print("QueryRow failed: ", err)
		json.NewEncoder(w).Encode(resp{Access: false})
	} else {
		log.Print("User entry")
		json.NewEncoder(w).Encode(resp{Access: true, Login: user.Login, ID: id_user})
	}

}

//Get information about room
func GetInfoAboutRooms(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	decoder := json.NewDecoder(r.Body)
	var roomID map[string]string
	err := decoder.Decode(&roomID)
	if err != nil {
		log.Print("JSON faild: ", err)
	}

	type resp struct {
		Access bool `json:"Access"`
	}

	checkRoom := AllRooms.Map[roomID["RoomID"]]

	if checkRoom == nil {
		json.NewEncoder(w).Encode(resp{Access: false})
	}else{
		json.NewEncoder(w).Encode(resp{Access: true})
	}
}

// Create a Room and return roomID
func CreateRoomReguestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomID := AllRooms.CreateRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}

	log.Print(AllRooms.Map)
	json.NewEncoder(w).Encode(resp{RoomID: roomID})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type broadcastMsg struct {
	Message map[string]interface{}
	RoomId  string
	Client  *websocket.Conn
}

var broadcast = make(chan broadcastMsg)
var connecshons uint8 = 0

func broadcaster() {
	for {
		msg := <-broadcast

		for _, client := range AllRooms.Map[msg.RoomId] {
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)

				if err != nil {
					log.Print(err, " In roomID ", msg.RoomId)
					client.Conn.Close()
					connecshons -= 1
					break
				}
			}
			if (client.Conn == msg.Client) && (connecshons == 0) {
				err := client.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Print(err)
					client.Conn.Close()
					AllRooms.DeleteRoom(msg.RoomId)
					break
				}
			}
		}
	}
}

// Join the client in a particular room
func JoinRoomReguestHandler(w http.ResponseWriter, r *http.Request) {

	roomID, ok := r.URL.Query()["roomID"]
	if !ok {
		log.Print("roomID missing in URL Parameters")
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
	}

	AllRooms.InsertIntoRoom(roomID[0], false, ws)

	connecshons += 1

	var goinAway uint8 = 0
	var readConnections = AllRooms.Map[roomID[0]]
	log.Printf("RoomID %s and Connecshons %v ", roomID, readConnections)

	go broadcaster()

	for {
		var msg broadcastMsg

		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Print("Read Error: ", err)
			goinAway += 1
			switch goinAway {
			case 20:
				continue
			case 21:
				break
			}
		}

		msg.Client = ws
		msg.RoomId = roomID[0]
		//log.Println(msg.Message)

		broadcast <- msg
	}
}
