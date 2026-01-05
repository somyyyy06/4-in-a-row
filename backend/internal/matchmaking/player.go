package matchmaking

import (
	"time"

	"github.com/gorilla/websocket"
)

type Player struct {
	Username string
	Conn     *websocket.Conn
	BotTimer *time.Timer
}
