package game

type Bot struct {
	Player int
}

func NewBot(player int) *Bot {
	return &Bot{Player: player}
}

func (b *Bot) ChooseMove(board Board, opponent int) int {

	for col := 0; col < Cols; col++ {
		if board.IsColumnFull(col) {
			continue
		}

		simulated := board
		row, _ := simulated.DropDisc(col, b.Player)

		if CheckWin(simulated, row, col, b.Player) {
			return col
		}
	}

	for col := 0; col < Cols; col++ {
		if board.IsColumnFull(col) {
			continue
		}

		simulated := board
		row, _ := simulated.DropDisc(col, opponent)

		if CheckWin(simulated, row, col, opponent) {
			return col
		}
	}

	preferred := []int{3, 2, 4, 1, 5, 0, 6}
	for _, col := range preferred {
		if !board.IsColumnFull(col) {
			return col
		}
	}

	return -1
}
