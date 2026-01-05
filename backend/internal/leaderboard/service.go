package leaderboard

import "fourinarow/internal/db"

func SaveGame(gameID, p1, p2, winner string) error {
	_, err := db.DB.Exec(
		"INSERT INTO games (id, player1, player2, winner) VALUES (?, ?, ?, ?)",
		gameID, p1, p2, winner,
	)
	return err
}

func IncrementWin(username string) error {
	_, err := db.DB.Exec(
		`INSERT INTO leaderboard (username, wins)
		 VALUES (?, 1)
		 ON DUPLICATE KEY UPDATE wins = wins + 1`,
		username,
	)
	return err
}
