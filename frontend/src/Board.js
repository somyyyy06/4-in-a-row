export default function Board({ board, onMove }) {
  return (
    <div style={styles.board}>
      {board.map((row, r) =>
        row.map((cell, c) => (
          <div
            key={`${r}-${c}`}
            onClick={() => onMove(c)}
            style={{
              ...styles.cell,
              backgroundColor:
                cell === 1 ? "#ef4444" :
                cell === 2 ? "#facc15" :
                "#e5e7eb"
            }}
          />
        ))
      )}
    </div>
  );
}

const styles = {
  board: {
    display: "grid",
    gridTemplateColumns: "repeat(7, 60px)",
    gap: "6px",
    backgroundColor: "#1f2937",
    padding: "10px",
    borderRadius: "10px",
    margin: "20px auto"
  },
  cell: {
    width: "60px",
    height: "60px",
    borderRadius: "50%",
    cursor: "pointer",
    transition: "transform 0.1s ease",
  }
};
