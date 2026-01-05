import { useEffect, useState } from "react";

export default function Leaderboard({ refresh }) {
  const [data, setData] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/leaderboard")
      .then(res => res.json())
      .then(result => {
        if (Array.isArray(result)) {
          setData(result);
        } else {
          setData([]);
        }
      })
      .catch(err => {
        console.error("Leaderboard fetch failed:", err);
        setData([]);
      });
  }, [refresh]); 

  return (
    <div className="card">
      <h2>🏆 Leaderboard</h2>

      {data.length === 0 ? (
        <p>No games played yet.</p>
      ) : (
        data.map((p, i) => (
          <div key={p.username} style={row}>
            <span style={rank}>#{i + 1}</span>
            <span style={{ flex: 1 }}>{p.username}</span>
            <strong>{p.wins}W</strong>
          </div>
        ))
      )}
    </div>
  );
}

const row = {
  display: "flex",
  alignItems: "center",
  padding: "12px",
  borderRadius: "10px",
  background: "#f8fafc",
  marginTop: "10px"
};

const rank = {
  width: "30px",
  fontWeight: "bold"
};
