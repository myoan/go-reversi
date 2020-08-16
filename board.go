package reversi

import (
	"fmt"
)

type Board struct {
	Width  int
	Height int
	board  [][]*Cell
}

func NewBoard(init [][]int) *Board {
	height := len(init)
	if height == 0 {
		return &Board{board: [][]*Cell{}}
	}

	width := len(init[0])
	board := make([][]*Cell, height)

	for i, line := range init {
		bLine := make([]*Cell, width)
		for j, state := range line {
			bLine[j] = &Cell{X: j, Y: i, State: state}
		}
		board[i] = bLine
	}
	return &Board{board: board, Width: width, Height: height}
}

func (b *Board) GetBoard() [][]*Cell {
	return b.board
}

func (b *Board) Cell(x, y int) *Cell {
	if x < 0 || x >= b.Width {
		return nil
	}
	if y < 0 || y >= b.Height {
		return nil
	}
	return b.board[y][x]
}

func (b *Board) Show() {
	fmt.Println("----------------------")
	fmt.Println("    0 1 2 3 4 5 6 7")
	for i := 0; i < b.Height; i++ {
		fmt.Printf(" %d  ", i)
		for j := 0; j < b.Width; j++ {
			cell := b.board[i][j]
			switch cell.State {
			case 0:
				fmt.Print("_")
			case 1:
				fmt.Print("x")
			case 2:
				fmt.Print("o")
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (b *Board) SetStone(color int, pos *Position) error {
	fmt.Printf("SetStone: (%d, %d) color: %d\n", pos.X, pos.Y, color)
	cell := b.Cell(pos.X, pos.Y)
	if cell == nil {
		return fmt.Errorf("Cell not found at (%d, %d)", pos.X, pos.Y)
	}
	if cell.State != int(None) {
		return fmt.Errorf("Not empty cell (%d, %d)", pos.X, pos.Y)
	}
	if err := b.allocate(color, cell); err != nil {
		return fmt.Errorf("Cell not allocate (%d, %d)", pos.X, pos.Y)
	}
	return nil
}

func (b *Board) IsOccupied() bool {
	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			c := b.Cell(i, j)
			if c.State == 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) Opponent(color int) int {
	if color == 1 {
		return 2
	}
	return 1
}

func (b *Board) CanAllocate(color int) bool {
	// opponent = b.opponent(color)
	return true
}

func (b *Board) toArray() [][]int {
	ret := make([][]int, 0, b.Height)
	for _, line := range b.board {
		a := make([]int, 0, b.Width)
		for _, cell := range line {
			a = append(a, cell.State)
		}
		ret = append(ret, a)
	}
	return ret
}

func (b *Board) allocate(color int, cell *Cell) error {
	var allocated = false
	var ds = []*Direction{
		{dx: 0, dy: -1},  // top
		{dx: 1, dy: -1},  // top right
		{dx: 1, dy: 0},   // right
		{dx: 1, dy: 1},   // bottom right
		{dx: 0, dy: 1},   // bottom
		{dx: -1, dy: 1},  // bottom left
		{dx: -1, dy: 0},  // left
		{dx: -1, dy: -1}, // top left
	}
	for _, d := range ds {
		if b.seek(d, color, cell) {
			b.update(d, color, cell)
			allocated = true
		}
	}
	if allocated {
		return nil
	}
	return fmt.Errorf("Failed to allocate at (%d, %d)", cell.X, cell.Y)
}

func (b *Board) next(d *Direction, cell *Cell) (*Cell, error) {
	x, y := d.Next(cell.X, cell.Y)
	if x < 0 || x >= b.Width {
		return nil, fmt.Errorf("Invalid position")
	}
	if y < 0 || y >= b.Height {
		return nil, fmt.Errorf("Invalid position")
	}
	return b.Cell(x, y), nil
}

func (b *Board) seek(d *Direction, color int, cell *Cell) bool {
	x, y := d.Next(cell.X, cell.Y)
	next := b.Cell(x, y)
	if next == nil {
		return false
	}

	opponent := b.Opponent(color)

	switch next.State {
	case opponent:
		return b.seek(d, color, next)
	case color:
		if cell.State == opponent {
			return true
		}
		return false
	default:
		return false
	}
}

func (b *Board) update(d *Direction, color int, cell *Cell) {
	x, y := d.Next(cell.X, cell.Y)
	next := b.Cell(x, y)
	if next == nil {
		return
	}
	cell.Update(color)
	opponent := b.Opponent(color)

	if next.State == opponent {
		b.update(d, color, next)
	}
}
