export default function GameBoard({ board = defaultBoard, onMove }) {
  return (
    <div style={boardWrapper}>
      {/* Column click overlay */}
      <div style={columnOverlay}>
        {Array(7).fill(0).map((_, col) => (
          <div
            key={col}
            style={columnHitbox}
            onClick={() => onMove && onMove(col)}
          />
        ))}
      </div>

      {/* Actual board */}
      <div style={boardStyle}>
        {board.map((row, r) =>
          row.map((cell, c) => (
            <div
              key={`${r}-${c}`}
              style={{
                ...cellStyle,
                background:
                  cell === 1 ? "#ef4444" :
                  cell === 2 ? "#facc15" :
                  "#e5e7eb"
              }}
            />
          ))
        )}
      </div>
    </div>
  );
}


const defaultBoard = Array(6)
  .fill(0)
  .map(() => Array(7).fill(0));


const boardWrapper = {
  position: "relative",
  width: "max-content",
  margin: "20px auto"
};

const boardStyle = {
  display: "grid",
  gridTemplateColumns: "repeat(7, 60px)",
  gap: "8px",
  background: "#1f2937",
  padding: "12px",
  borderRadius: "12px"
};

const cellStyle = {
  width: "60px",
  height: "60px",
  borderRadius: "50%",
  transition: "background 0.2s"
};

const columnOverlay = {
  position: "absolute",
  top: 0,
  left: 0,
  display: "grid",
  gridTemplateColumns: "repeat(7, 60px)",
  gap: "8px",
  height: "100%",
  width: "100%",
  zIndex: 10
};

const columnHitbox = {
  cursor: "pointer"
};
