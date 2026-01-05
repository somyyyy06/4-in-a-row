package websocket

import (
	"log"
	"time"

	"fourinarow/internal/matchmaking"
	"fourinarow/internal/reconnect"

	"github.com/gorilla/websocket"
)

func handleConnection(conn *websocket.Conn) {

	var joinMsg struct {
		Type     string `json:"type"`
		Username string `json:"username"`
	}

	if err := conn.ReadJSON(&joinMsg); err != nil {
		log.Println("Failed to read join message:", err)
		return
	}

	if joinMsg.Type != "join" || joinMsg.Username == "" {
		log.Println("Invalid join payload")
		return
	}

	username := joinMsg.Username
	log.Println("Player joined:", username)

	if gameSession, ok := reconnect.CanReconnect(username); ok {
		log.Println("Reconnecting player:", username)
		attachPlayer(gameSession, username, conn)
		return
	}

	player := &matchmaking.Player{
		Username: username,
		Conn:     conn,
	}

	opponent, matched := MatchQueue.AddPlayer(player)

	if matched {
		log.Println("Match found:", player.Username, "vs", opponent.Username)

		if player.BotTimer != nil {
			player.BotTimer.Stop()
		}
		if opponent.BotTimer != nil {
			opponent.BotTimer.Stop()
		}

		startPlayerGame(player, opponent)
		return
	}

	log.Println("Waiting for opponent:", player.Username)

	matchmaking.StartBotTimer(player, 10*time.Second, func() {
		log.Println("Bot timeout reached for:", player.Username)

		MatchQueue.RemovePlayer(player)

		if player.Conn != nil {
			startBotGame(player)
		}
	})
}
