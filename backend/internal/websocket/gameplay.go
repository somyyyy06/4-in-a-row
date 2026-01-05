package websocket

import (
	"log"

	"fourinarow/internal/game"
	"fourinarow/internal/leaderboard"
)


func gameplayLoop(session *game.GameSession) {
	players := []*game.PlayerSession{
		session.Player1,
		session.Player2,
	}

	for {
		current := players[session.Engine.Turn-1]

		var msg struct {
			Type   string `json:"type"`
			Column int    `json:"column"`
		}

		if err := current.Conn.ReadJSON(&msg); err != nil {
			log.Println("Disconnect detected:", current.Username)
			return
		}

		if msg.Type != "move" {
			continue
		}

		win, draw, err := session.Engine.MakeMove(msg.Column)
		if err != nil {
			continue
		}

		for _, p := range players {
			sendUpdate(p.Conn, session.Engine.Board, session.Engine.Turn)
		}

		if win {
			if session.Over {
				return
			}
			session.Over = true

			winner := current.Username

			leaderboard.SaveGame(
				session.ID,
				session.Player1.Username,
				session.Player2.Username,
				winner,
			)
			leaderboard.IncrementWin(winner)

			for _, p := range players {
				result := "loss"
				if p.Username == winner {
					result = "win"
				}
				sendEnd(p.Conn, result, winner)
			}
			return
		}

		if draw {
			if session.Over {
				return
			}
			session.Over = true

			leaderboard.SaveGame(
				session.ID,
				session.Player1.Username,
				session.Player2.Username,
				"draw",
			)

			for _, p := range players {
				sendEnd(p.Conn, "draw", "draw")
			}
			return
		}
	}
}


func gameplayLoopWithBot(session *game.GameSession, bot *game.Bot) {
	player := session.Player1

	for {
		var msg struct {
			Type   string `json:"type"`
			Column int    `json:"column"`
		}

		if err := player.Conn.ReadJSON(&msg); err != nil {
			log.Println("Player disconnected")
			return
		}

		if msg.Type != "move" {
			continue
		}

		win, draw, err := session.Engine.MakeMove(msg.Column)
		if err != nil {
			continue
		}

		sendUpdate(player.Conn, session.Engine.Board, session.Engine.Turn)

		if win {
			if session.Over {
				return
			}
			session.Over = true

			winner := player.Username

			leaderboard.SaveGame(
				session.ID,
				player.Username,
				"BOT",
				winner,
			)
			leaderboard.IncrementWin(winner)

			sendEnd(player.Conn, "win", winner)
			return
		}

		if draw {
			if session.Over {
				return
			}
			session.Over = true

			leaderboard.SaveGame(
				session.ID,
				player.Username,
				"BOT",
				"draw",
			)

			sendEnd(player.Conn, "draw", "draw")
			return
		}

		botCol := bot.ChooseMove(session.Engine.Board, game.Player1)
		win, draw, _ = session.Engine.MakeMove(botCol)

		sendUpdate(player.Conn, session.Engine.Board, session.Engine.Turn)

		if win {
			if session.Over {
				return
			}
			session.Over = true

			leaderboard.SaveGame(
				session.ID,
				player.Username,
				"BOT",
				"BOT",
			)

			sendEnd(player.Conn, "loss", "BOT")
			return
		}

		if draw {
			if session.Over {
				return
			}
			session.Over = true

			leaderboard.SaveGame(
				session.ID,
				player.Username,
				"BOT",
				"draw",
			)

			sendEnd(player.Conn, "draw", "draw")
			return
		}
	}
}
