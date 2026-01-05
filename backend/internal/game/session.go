package game

import (
	"time"

	"github.com/gorilla/websocket"
)

type PlayerSession struct {
	Username string
	Conn     *websocket.Conn
}

type GameSession struct {
	ID           string
	Engine       *GameEngine
	Player1      *PlayerSession
	Player2      *PlayerSession
	Disconnected map[string]time.Time
	Over         bool
}
