import { useEffect, useState } from "react";

export default function Leaderboard() {
  const [data, setData] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/leaderboard")
      .then(res => res.json())
      .then(setData);
  }, []);

  return (
    <div style={styles.container}>
      <h2>🏆 Leaderboard</h2>
      <ul style={styles.list}>
        {data.map(row => (
          <li key={row.username} style={styles.item}>
            <span>{row.username}</span>
            <strong>{row.wins}</strong>
          </li>
        ))}
      </ul>
    </div>
  );
}

const styles = {
  container: {
    maxWidth: "300px",
    margin: "30px auto",
    background: "white",
    padding: "20px",
    borderRadius: "10px",
    boxShadow: "0 4px 12px rgba(0,0,0,0.1)"
  },
  list: {
    listStyle: "none",
    padding: 0,
    margin: 0
  },
  item: {
    display: "flex",
    justifyContent: "space-between",
    padding: "8px 0",
    borderBottom: "1px solid #e5e7eb"
  }
};
