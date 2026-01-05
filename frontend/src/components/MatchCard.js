export default function MatchCard({
  username,
  setUsername,
  started,
  onJoin,
  children
}) {
  return (
    <div className="card" style={{ textAlign: "center" }}>
      {!started ? (
        <>
          <div style={icon}>⚡</div>
          <h2>Ready for a Challenge?</h2>
          <p>
            Join the queue to be matched with a player or our smart bot.
            Games are fast, fun, and real-time.
          </p>

          {/* USERNAME INPUT — ALWAYS VISIBLE */}
          <input
            placeholder="Enter username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            style={{
              padding: "12px",
              width: "70%",
              marginBottom: "15px",
              borderRadius: "8px",
              border: "1px solid #ccc"
            }}
          />

          <br />

          <button className="primary" onClick={onJoin}>
            Find Match
          </button>
        </>
      ) : (
        children
      )}
    </div>
  );
}

const icon = {
  width: "70px",
  height: "70px",
  borderRadius: "50%",
  background: "#fff2e5",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  fontSize: "28px",
  margin: "0 auto 20px"
};
