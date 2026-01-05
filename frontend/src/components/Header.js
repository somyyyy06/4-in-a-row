export default function Header({ username }) {
  return (
    <div style={styles.header}>
      <div>
        <h1>4-in-a-Row</h1>
        <p>Connect four discs to win. Simple as that.</p>
      </div>

      <div style={styles.user}>
        <span style={styles.dot}></span>
        {username}
        <span style={styles.logout}>Log Out</span>
      </div>
    </div>
  );
}

const styles = {
  header: {
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center"
  },
  user: {
    display: "flex",
    alignItems: "center",
    gap: "10px",
    background: "#fff",
    padding: "10px 16px",
    borderRadius: "20px"
  },
  dot: {
    width: "8px",
    height: "8px",
    background: "#22c55e",
    borderRadius: "50%"
  },
  logout: {
    marginLeft: "10px",
    color: "#64748b",
    cursor: "pointer"
  }
};
