package game

import "errors"

const (
	Rows = 6
	Cols = 7

	Empty   = 0
	Player1 = 1
	Player2 = 2
)

type Board [Rows][Cols]int

func NewBoard() Board {
	var b Board
	return b
}

func (b *Board) DropDisc(col int, player int) (int, error) {
	if col < 0 || col >= Cols {
		return -1, errors.New("invalid column")
	}

	for row := Rows - 1; row >= 0; row-- {
		if b[row][col] == Empty {
			b[row][col] = player
			return row, nil
		}
	}

	return -1, errors.New("column is full")
}

func (b Board) IsColumnFull(col int) bool {
	return b[0][col] != Empty
}
