package main

import (
	"log"
	"net/http"
	"web_conf_0.2/server"
)

func main() {
	server.AllRooms.Init()

	http.HandleFunc("/create", server.CreateRoomReguestHandler)
	http.HandleFunc("/join", server.JoinRoomReguestHandler)
	http.HandleFunc("/login", server.LoginReguestHandler)
	http.HandleFunc("/GETInfRooms", server.GetInfoAboutRooms)

	log.Print("Starting in port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
