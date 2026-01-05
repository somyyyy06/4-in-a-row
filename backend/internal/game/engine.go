package game

type GameEngine struct {
	Board Board
	Turn  int
}

func NewGameEngine() *GameEngine {
	return &GameEngine{
		Board: NewBoard(),
		Turn:  Player1,
	}
}

func (g *GameEngine) MakeMove(col int) (win bool, draw bool, err error) {
	row, err := g.Board.DropDisc(col, g.Turn)
	if err != nil {
		return false, false, err
	}

	win = CheckWin(g.Board, row, col, g.Turn)
	if !win {
		draw = IsDraw(g.Board)
	}

	if g.Turn == Player1 {
		g.Turn = Player2
	} else {
		g.Turn = Player1
	}

	return win, draw, nil
}
