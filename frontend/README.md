# 🎮 4-in-a-Row (Connect Four) — Real-Time Multiplayer Game

A real-time, backend-driven implementation of the classic **4-in-a-Row (Connect Four)** game, built as part of a backend engineering assignment.

The game supports:
- Real-time multiplayer using WebSockets
- Automatic bot fallback if no opponent joins
- Persistent leaderboard
- Reconnection handling
- Clean React frontend
- Go backend with MySQL persistence

---

## 🚀 Features

- 🎯 **Real-time gameplay** using WebSockets
- 👥 **Player vs Player** matchmaking
- 🤖 **Competitive bot fallback** after timeout
- 🔁 **Reconnect support** for disconnected players
- 🏆 **Live leaderboard** with persistent stats
- 🗄️ **MySQL-backed game storage**
- ⚙️ **Clean backend architecture (Go)**
- 🖥️ **Simple, functional React UI**

---

## 🧠 Tech Stack

### Backend
- **Go**
- Gorilla WebSocket
- MySQL
- Clean modular architecture

### Frontend
- **React**
- WebSocket API
- Basic responsive UI

### Database
- **MySQL**


---

## ⚙️ Prerequisites

Make sure you have the following installed:

- **Go** (>= 1.20)
- **Node.js** (>= 16)
- **npm**
- **MySQL** (>= 8.0)
- **Git**

---

## 🗄️ Database Setup (MySQL)

### 1️⃣ Create Database
```sql
CREATE DATABASE fourinarow;
USE fourinarow;


2️⃣ Create Tables
CREATE TABLE games (
  id VARCHAR(255) PRIMARY KEY,
  player1 VARCHAR(100),
  player2 VARCHAR(100),
  winner VARCHAR(100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE leaderboard (
  username VARCHAR(100) PRIMARY KEY,
  wins INT DEFAULT 0
);


3️⃣ Configure Database Connection

Update MySQL credentials in:

backend/cmd/server/main.go


Example:

db.InitMySQL("root:yourpassword@tcp(localhost:3306)/fourinarow")

▶️ Running the Project Locally
🔹 Backend (Go)

From the project root:

cd backend
go mod tidy
go run cmd/server/main.go


You should see logs similar to:

MySQL connected
4-in-a-Row server started on :8080

🔹 Frontend (React)

Open a new terminal window:

cd frontend
npm install
npm start


The frontend will be available at:

http://localhost:3000

🎮 How to Play
Player vs Player

Open two browser tabs (or one normal + one incognito)

Enter different usernames in each tab

Click Find Match in both tabs

Play the game in real time

Player vs Bot

Open one browser tab

Enter a username

Click Find Match

Wait ~10 seconds

A competitive bot joins automatically

🏆 Leaderboard

Updates in real time after every game

Stores total wins per player

Data is persisted in MySQL

No page refresh required

🔁 Reconnection Handling

If a player disconnects mid-game, they can reconnect using the same username

Reconnection window: 30 seconds

If the player does not reconnect, the opponent (or bot) wins by forfeit
