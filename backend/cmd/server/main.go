package main

import (
	"log"
	"net/http"
	"os"

	"fourinarow/internal/db"
	ws "fourinarow/internal/websocket"
)

func main() {
	mysqlDSN := os.Getenv("MYSQL_DSN")
	if mysqlDSN == "" {
		log.Fatal("MYSQL_DSN environment variable not set")
	}
	db.InitMySQL(mysqlDSN)

	http.HandleFunc("/ws", ws.HandleWebSocket)
	http.HandleFunc("/leaderboard", ws.LeaderboardHandler)

	handler := ws.EnableCORS(http.DefaultServeMux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("4-in-a-Row server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
