package matchmaking

import (
	"sync"
	"time"
)

type Queue struct {
	players []*Player
	mu      sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) AddPlayer(p *Player) (*Player, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.players) > 0 {
		opponent := q.players[0]
		q.players = q.players[1:]
		return opponent, true
	}

	q.players = append(q.players, p)
	return nil, false
}

func (q *Queue) RemovePlayer(p *Player) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for i, player := range q.players {
		if player == p {
			q.players = append(q.players[:i], q.players[i+1:]...)
			break
		}
	}
}

func StartBotTimer(p *Player, timeout time.Duration, onTimeout func()) {
	p.BotTimer = time.AfterFunc(timeout, onTimeout)
}
