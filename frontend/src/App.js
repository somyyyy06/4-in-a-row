import { useEffect, useRef, useState } from "react";
import Header from "./components/Header";
import MatchCard from "./components/MatchCard";
import GameBoard from "./components/GameBoard";
import Leaderboard from "./components/Leaderboard";

const emptyBoard = Array(6)
  .fill(0)
  .map(() => Array(7).fill(0));

export default function App() {
  const socketRef = useRef(null);

  const [username, setUsername] = useState("");
  const [started, setStarted] = useState(false);
  const [player, setPlayer] = useState(null);
  const [turn, setTurn] = useState(null);
  const [board, setBoard] = useState(emptyBoard);
  const [status, setStatus] = useState("");
  const [leaderboardRefresh, setLeaderboardRefresh] = useState(0);
  const [gameOver, setGameOver] = useState(false);


  function findMatch() {
    if (!username.trim()) {
      alert("Please enter a username first");
      return;
    }

    if (socketRef.current) return;

    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => {
      ws.send(JSON.stringify({ type: "join", username }));
    };

    ws.onmessage = (event) => {
      const msg = JSON.parse(event.data);

      if (msg.type === "start") {
        setPlayer(msg.player);
        setTurn(1);
        setStarted(true);
        setGameOver(false);
        setStatus("");
        setBoard(emptyBoard);
      }

      if (msg.type === "update") {
        setBoard(msg.board);
        setTurn(msg.turn);
      }

      if (msg.type === "end") {
        setGameOver(true);

        if (msg.result === "win") {
          setStatus("🎉 You won!");
        } else if (msg.result === "loss") {
          setStatus(`❌ You lost. Winner: ${msg.winner}`);
        } else {
          setStatus("🤝 Draw");
        }

        setLeaderboardRefresh(Date.now());
      }
    };

    ws.onclose = () => {
      socketRef.current = null;
    };

    socketRef.current = ws;
  }


  function makeMove(column) {
    if (!started) return;
    if (gameOver) return;
    if (player !== turn) return;

    socketRef.current.send(
      JSON.stringify({ type: "move", column })
    );
  }


  useEffect(() => {
    const handleUnload = () => {
      if (socketRef.current) {
        socketRef.current.close();
        socketRef.current = null;
      }
    };

    window.addEventListener("beforeunload", handleUnload);

    return () => {
      window.removeEventListener("beforeunload", handleUnload);
      handleUnload();
    };
  }, []);


  return (
    <div className="container">
      <Header username={username || "Guest"} />

      <div className="main-grid">
        <MatchCard
          username={username}
          setUsername={setUsername}
          started={started}
          onJoin={findMatch}
        >
          <>
            <GameBoard board={board} onMove={makeMove} />
            <h3 style={{ marginTop: "10px" }}>
              {status ||
                (player === turn
                  ? "🟢 Your turn"
                  : "🟡 Opponent's turn")}
            </h3>
          </>
        </MatchCard>

        <Leaderboard refresh={leaderboardRefresh} />
      </div>
    </div>
  );
}
