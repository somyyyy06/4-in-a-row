package websocket

import "github.com/gorilla/websocket"

func sendStart(conn *websocket.Conn, player int) {
	conn.WriteJSON(map[string]interface{}{
		"type":   "start",
		"player": player,
	})
}

func sendUpdate(conn *websocket.Conn, board interface{}, turn int) {
	conn.WriteJSON(map[string]interface{}{
		"type":  "update",
		"board": board,
		"turn":  turn,
	})
}

func sendEnd(conn *websocket.Conn, result string, winner string) {
	conn.WriteJSON(map[string]interface{}{
		"type":   "end",
		"result": result,
		"winner": winner,
	})
}
