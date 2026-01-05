package game

import "testing"

func TestDropDisc(t *testing.T) {
	board := NewBoard()

	row, err := board.DropDisc(3, Player1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if row != Rows-1 {
		t.Fatalf("expected row %d, got %d", Rows-1, row)
	}

	if board[row][3] != Player1 {
		t.Fatalf("disc not placed correctly")
	}
}

func TestVerticalWin(t *testing.T) {
	engine := NewGameEngine()

	engine.MakeMove(0) 
	engine.MakeMove(1) 
	engine.MakeMove(0) 
	engine.MakeMove(1) 
	engine.MakeMove(0) 
	engine.MakeMove(1) 
	win, _, _ := engine.MakeMove(0) 

	if !win {
		t.Fatal("expected vertical win")
	}
}

func TestHorizontalWin(t *testing.T) {
	engine := NewGameEngine()

	engine.MakeMove(0)
	engine.MakeMove(0)
	engine.MakeMove(1)
	engine.MakeMove(1)
	engine.MakeMove(2)
	engine.MakeMove(2)
	win, _, _ := engine.MakeMove(3)

	if !win {
		t.Fatal("expected horizontal win")
	}
}

func TestDiagonalWin(t *testing.T) {
	engine := NewGameEngine()

	engine.MakeMove(0)
	engine.MakeMove(1)
	engine.MakeMove(1)
	engine.MakeMove(2)
	engine.MakeMove(2)
	engine.MakeMove(3)
	engine.MakeMove(2)
	engine.MakeMove(3)
	engine.MakeMove(3)
	engine.MakeMove(6)
	win, _, _ := engine.MakeMove(3)

	if !win {
		t.Fatal("expected diagonal win")
	}
}

func TestDraw(t *testing.T) {
	board := NewBoard()

	player := Player1
	for col := 0; col < Cols; col++ {
		for row := 0; row < Rows; row++ {
			board.DropDisc(col, player)
			if player == Player1 {
				player = Player2
			} else {
				player = Player1
			}
		}
	}

	if !IsDraw(board) {
		t.Fatal("expected draw")
	}
}

func TestBotBlocksWin(t *testing.T) {
	board := NewBoard()
	bot := NewBot(Player2)

	board.DropDisc(0, Player1)
	board.DropDisc(1, Player1)
	board.DropDisc(2, Player1)

	move := bot.ChooseMove(board, Player1)
	if move != 3 {
		t.Fatalf("expected bot to block at column 3, got %d", move)
	}
}
