let socket = null;

export function connect(username, onMessage) {
  socket = new WebSocket("ws://localhost:8080/ws");

  socket.onopen = () => {
    socket.send(JSON.stringify({
      type: "join",
      username
    }));
  };

  socket.onmessage = (event) => {
    onMessage(JSON.parse(event.data));
  };
}

export function sendMove(column) {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify({
      type: "move",
      column
    }));
  }
}
