package websocket

import (
	"encoding/json"
	"net/http"

	"fourinarow/internal/db"
)

func LeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT username, wins FROM leaderboard ORDER BY wins DESC")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type entry struct {
		Username string `json:"username"`
		Wins     int    `json:"wins"`
	}

	var result []entry
	for rows.Next() {
		var e entry
		rows.Scan(&e.Username, &e.Wins)
		result = append(result, e)
	}

	json.NewEncoder(w).Encode(result)
}
