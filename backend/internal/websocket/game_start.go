package websocket
import "log"

import (
	"time"

	"fourinarow/internal/game"
	"fourinarow/internal/matchmaking"
	"fourinarow/internal/reconnect"
)

func startPlayerGame(p1, p2 interface{}) {
	player1 := p1.(*matchmaking.Player)
	player2 := p2.(*matchmaking.Player)

	log.Println("Starting PvP game between", player1.Username, "and", player2.Username)

	engine := game.NewGameEngine()

	session := &game.GameSession{
		ID:           player1.Username + "_vs_" + player2.Username,
		Engine:       engine,
		Player1:      &game.PlayerSession{Username: player1.Username, Conn: player1.Conn},
		Player2:      &game.PlayerSession{Username: player2.Username, Conn: player2.Conn},
		Disconnected: make(map[string]time.Time),
		Over:         false,
	}

	reconnect.RegisterGame(session)

	sendStart(player1.Conn, game.Player1)
	sendStart(player2.Conn, game.Player2)

	go gameplayLoop(session)
}

func startBotGame(p interface{}) {
	player := p.(*matchmaking.Player)

	engine := game.NewGameEngine()
	bot := game.NewBot(game.Player2)

	session := &game.GameSession{
		ID:           player.Username + "_vs_bot",
		Engine:       engine,
		Player1:      &game.PlayerSession{Username: player.Username, Conn: player.Conn},
		Player2:      &game.PlayerSession{Username: "BOT", Conn: nil},
		Disconnected: make(map[string]time.Time),
		Over:         false,
	}

	reconnect.RegisterGame(session)

	sendStart(player.Conn, game.Player1)

	go gameplayLoopWithBot(session, bot)
}
