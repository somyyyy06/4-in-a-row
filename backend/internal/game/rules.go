package game

func countInDirection(b Board, row, col, dr, dc, player int) int {
	count := 0
	r := row + dr
	c := col + dc

	for r >= 0 && r < Rows && c >= 0 && c < Cols && b[r][c] == player {
		count++
		r += dr
		c += dc
	}

	return count
}

func CheckWin(b Board, row, col, player int) bool {
	directions := [][]int{
		{0, 1},   
		{1, 0},   
		{1, 1},   
		{-1, 1},  
	}

	for _, d := range directions {
		count := 1
		count += countInDirection(b, row, col, d[0], d[1], player)
		count += countInDirection(b, row, col, -d[0], -d[1], player)

		if count >= 4 {
			return true
		}
	}

	return false
}

func IsDraw(b Board) bool {
	for col := 0; col < Cols; col++ {
		if b[0][col] == Empty {
			return false
		}
	}
	return true
}
