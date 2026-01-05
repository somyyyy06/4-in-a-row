package reconnect

import (
	"time"

	"fourinarow/internal/game"
)

func StartForfeitTimer(game *game.GameSession, username string, onForfeit func()) {
	go func() {
		time.Sleep(30 * time.Second)

		if _, stillDisconnected := game.Disconnected[username]; stillDisconnected {
			onForfeit()
		}
	}()
}
