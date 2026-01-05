package reconnect

import (
	"sync"
	"time"

	"fourinarow/internal/game"
)

var (
	GamesByID       = make(map[string]*game.GameSession)
	GamesByUsername = make(map[string]string)
	mu              sync.Mutex
)

func RegisterGame(game *game.GameSession) {
	mu.Lock()
	defer mu.Unlock()

	GamesByID[game.ID] = game
	GamesByUsername[game.Player1.Username] = game.ID
	GamesByUsername[game.Player2.Username] = game.ID
}

func MarkDisconnected(username string) {
	mu.Lock()
	defer mu.Unlock()

	gameID, ok := GamesByUsername[username]
	if !ok {
		return
	}

	game := GamesByID[gameID]
	game.Disconnected[username] = time.Now()
}

func CanReconnect(username string) (*game.GameSession, bool) {
	mu.Lock()
	defer mu.Unlock()

	gameID, ok := GamesByUsername[username]
	if !ok {
		return nil, false
	}

	game := GamesByID[gameID]
	t, disconnected := game.Disconnected[username]
	if !disconnected {
		return game, true
	}

	if time.Since(t) <= 30*time.Second {
		delete(game.Disconnected, username)
		return game, true
	}

	return nil, false
}
