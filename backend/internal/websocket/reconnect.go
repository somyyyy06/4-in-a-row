package websocket

import (
	"log"

	"fourinarow/internal/game"

	"github.com/gorilla/websocket"
)

func attachPlayer(gameSession *game.GameSession, username string, conn *websocket.Conn) {
	log.Println("Player reattached:", username)

	if gameSession.Player1.Username == username {
		gameSession.Player1.Conn = conn
	} else if gameSession.Player2.Username == username {
		gameSession.Player2.Conn = conn
	}

}
