package main

import (
	"log"
	"net/http"

	"fourinarow/internal/db"
	ws "fourinarow/internal/websocket"
)

func main() {
	db.InitMySQL("root:jatinsomya06080301@@tcp(localhost:3306)/fourinarow")

	http.HandleFunc("/ws", ws.HandleWebSocket)
	http.HandleFunc("/leaderboard", ws.LeaderboardHandler)

	handler := ws.EnableCORS(http.DefaultServeMux)

	log.Println("4-in-a-Row server started on :8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
