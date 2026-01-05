import { useEffect, useRef, useState } from "react";
import Board from "./Board";

export default function Game() {
  const socketRef = useRef(null);

  const [username, setUsername] = useState("");
  const [connected, setConnected] = useState(false);
  const [started, setStarted] = useState(false);
  const [player, setPlayer] = useState(null);
  const [turn, setTurn] = useState(null);
  const [board, setBoard] = useState(
    Array(6).fill(0).map(() => Array(7).fill(0))
  );
  const [status, setStatus] = useState("");

  function connectOnce() {
    if (socketRef.current) return;

    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => {
      ws.send(JSON.stringify({ type: "join", username }));
      setConnected(true);
    };

    ws.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      console.log("Server:", msg);

      if (msg.type === "start") {
        setPlayer(msg.player);
        setTurn(1);
        setStarted(true);
      }

      if (msg.type === "update") {
        setBoard(msg.board);
        setTurn(msg.turn);
      }

      if (msg.type === "end") {
        setStatus(msg.result.toUpperCase());
      }
    };

    ws.onclose = () => {
      console.log("WebSocket closed");
    };

    socketRef.current = ws;
  }

  useEffect(() => {
    return () => {
      if (socketRef.current) {
        socketRef.current.close();
        socketRef.current = null;
      }
    };
  }, []);

  function handleJoin() {
    if (!username) return;
    connectOnce();
  }

  function handleMove(column) {
    if (!started) return;
    if (player !== turn) return;

    socketRef.current.send(
      JSON.stringify({ type: "move", column })
    );
  }


  if (!connected) {
    return (
      <div>
        <input
          placeholder="Username"
          value={username}
          onChange={e => setUsername(e.target.value)}
        />
        <button onClick={handleJoin}>Join Game</button>
      </div>
    );
  }

  if (connected && !started) {
    return <h3>Waiting for opponent / bot...</h3>;
  }

  return (
  <div style={styles.container}>
    <Board board={board} onMove={handleMove} />
    <h3 style={{ marginTop: "10px" }}>
      {status ||
        (player === turn ? "🟢 Your turn" : "🟡 Opponent's turn")}
    </h3>
  </div>
);
}


const styles = {
  container: {
    maxWidth: "500px",
    margin: "0 auto",
    textAlign: "center",
  }
};
