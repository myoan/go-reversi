package reversi

import (
	"errors"
	"fmt"
)

type Game struct {
	GameState GameState
	board     *Board
}

type Position struct {
	X int
	Y int
}

type GameState int

const (
	Prepare GameState = iota
	BlackTurn
	WhiteTurn
	Finish
)

var InitBoard = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 2, 0, 0, 0},
	{0, 0, 0, 2, 1, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
}

func NewGame() *Game {
	board := NewBoard(InitBoard)
	return &Game{board: board}
}

func (game *Game) Show() {
	game.board.Show()
}

func (game *Game) GetBoard() [][]*Cell {
	return game.board.GetBoard()
}

func (game *Game) IsOccupied() bool {
	return game.board.IsOccupied()
}

func (game *Game) SetStone(color int, pos *Position) error {
	if game.GameState != GameState(color) {
		fmt.Printf("OutOfTurn: client: %d, server: %d\n", color, game.GameState)
		return errors.New("OutOfTurn")
	}

	err := game.board.SetStone(color, pos)
	if err != nil {
		return err
	}
	if game.board.IsOccupied() {
		game.updateGameState(Finish)
		return nil
	}
	opponent := game.board.Opponent(color)
	if len(game.board.ListAllocatablePositions(opponent)) > 0 {
		game.updateGameState(GameState(opponent))
	}
	return nil
}

func (game *Game) Winner() int {
	bCount := game.board.Count(int(Black))
	wCount := game.board.Count(int(White))
	if bCount > wCount {
		return 1
	} else if bCount < wCount {
		return 2
	} else {
		return 0
	}
}

func (game *Game) ListAllocatablePositions(color int) []*Position {
	return game.board.ListAllocatablePositions(color)
}

func (game *Game) updateGameState(s GameState) {
	fmt.Printf("set phase: %d -> %d\n", game.GameState, s)
	game.GameState = s
}
