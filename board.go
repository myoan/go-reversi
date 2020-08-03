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
	if x >= b.Width {
		return nil
	}
	if y >= b.Height {
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
	cell := b.board[pos.Y][pos.X]
	if cell == nil {
		return fmt.Errorf("Cell not found at (%d, %d)", pos.X, pos.Y)
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
	var opponent = b.Opponent(color)
	t, err := b.top(cell)
	if err == nil && t.State == opponent && b.seekTop(color, t) {
		b.updateTop(color, cell)
		allocated = true
	}
	tr, err := b.topRight(cell)
	if err == nil && tr.State == opponent && b.seekTopRight(color, tr) {
		b.updateTopRight(color, cell)
		allocated = true
	}
	r, err := b.right(cell)
	if err == nil && r.State == opponent && b.seekRight(color, r) {
		b.updateRight(color, cell)
		allocated = true
	}
	br, err := b.bottomRight(cell)
	if err == nil && br.State == opponent && b.seekBottomRight(color, br) {
		b.updateBottomRight(color, cell)
		allocated = true
	}
	btm, err := b.bottom(cell)
	if err == nil && btm.State == opponent && b.seekBottom(color, btm) {
		b.updateBottom(color, cell)
		allocated = true
	}
	bl, err := b.bottomLeft(cell)
	if err == nil && bl.State == opponent && b.seekBottomLeft(color, bl) {
		b.updateBottomLeft(color, cell)
		allocated = true
	}
	l, err := b.left(cell)
	if err == nil && l.State == opponent && b.seekLeft(color, l) {
		b.updateLeft(color, cell)
		allocated = true
	}
	tl, err := b.topLeft(cell)
	if err == nil && tl.State == opponent && b.seekTopLeft(color, tl) {
		b.updateTopLeft(color, cell)
		allocated = true
	}
	if allocated {
		return nil
	}
	return fmt.Errorf("Failed to allocate at (%d, %d)", cell.X, cell.Y)
}

func (b *Board) move(cell *Cell, nextX func(int) int, nextY func(int) int) (*Cell, error) {
	x := nextX(cell.X)
	y := nextY(cell.Y)
	if x < 0 || x >= b.Width {
		return nil, fmt.Errorf("Invalid position")
	}
	if y < 0 || y >= b.Height {
		return nil, fmt.Errorf("Invalid position")
	}
	return b.Cell(x, y), nil
}

func (b *Board) top(cell *Cell) (*Cell, error) {
	nextX := func(x int) int { return x }
	nextY := func(y int) int { return y - 1 }
	return b.move(cell, nextX, nextY)
}

func (b *Board) topRight(cell *Cell) (*Cell, error) {
	nextX := func(x int) int { return x + 1 }
	nextY := func(y int) int { return y - 1 }
	return b.move(cell, nextX, nextY)
}

func (b *Board) right(cell *Cell) (*Cell, error) {
	nextX := func(x int) int { return x + 1 }
	nextY := func(y int) int { return y }
	return b.move(cell, nextX, nextY)
}

func (b *Board) bottomRight(cell *Cell) (*Cell, error) {
	nextX := func(x int) int { return x + 1 }
	nextY := func(y int) int { return y + 1 }
	return b.move(cell, nextX, nextY)
}

func (b *Board) bottom(cell *Cell) (*Cell, error) {
	nextX := func(x int) int { return x }
	nextY := func(y int) int { return y + 1 }
	return b.move(cell, nextX, nextY)
}

func (b *Board) bottomLeft(cell *Cell) (*Cell, error) {
	nextX := func(x int) int { return x - 1 }
	nextY := func(y int) int { return y + 1 }
	return b.move(cell, nextX, nextY)
}

func (b *Board) left(cell *Cell) (*Cell, error) {
	nextX := func(x int) int { return x - 1 }
	nextY := func(y int) int { return y }
	return b.move(cell, nextX, nextY)
}

func (b *Board) topLeft(cell *Cell) (*Cell, error) {
	nextX := func(x int) int { return x - 1 }
	nextY := func(y int) int { return y - 1 }
	return b.move(cell, nextX, nextY)
}

func (b *Board) seekTop(color int, cell *Cell) bool {
	next, err := b.top(cell)
	if err != nil {
		return false
	}
	opponent := b.Opponent(color)

	switch next.State {
	case opponent:
		return b.seekTop(color, next)
	case color:
		if cell.State == opponent {
			return true
		}
		return false
	default:
		return false
	}
}

func (b *Board) updateTop(color int, cell *Cell) {
	next, err := b.top(cell)
	if err != nil {
		return
	}
	cell.Update(color)
	opponent := b.Opponent(color)

	if next.State == opponent {
		b.updateTop(color, next)
	}
}

func (b *Board) seekTopRight(color int, cell *Cell) bool {
	next, err := b.topRight(cell)
	if err != nil {
		return false
	}
	opponent := b.Opponent(color)

	switch next.State {
	case opponent:
		return b.seekTopRight(color, next)
	case color:
		if cell.State == opponent {
			return true
		}
		return false
	default:
		return false
	}
}

func (b *Board) updateTopRight(color int, cell *Cell) {
	next, err := b.topRight(cell)
	if err != nil {
		return
	}
	cell.Update(color)
	opponent := b.Opponent(color)

	if next.State == opponent {
		b.updateTopRight(color, next)
	}
}

func (b *Board) seekRight(color int, cell *Cell) bool {
	next, err := b.right(cell)
	if err != nil {
		return false
	}
	opponent := b.Opponent(color)

	switch next.State {
	case opponent:
		return b.seekRight(color, next)
	case color:
		if cell.State == opponent {
			return true
		}
		return false
	default:
		return false
	}
	return false
}

func (b *Board) updateRight(color int, cell *Cell) {
	next, err := b.right(cell)
	if err != nil {
		return
	}
	cell.Update(color)
	opponent := b.Opponent(color)

	if next.State == opponent {
		b.updateRight(color, next)
	}
}

func (b *Board) seekBottomRight(color int, cell *Cell) bool {
	next, err := b.bottomRight(cell)
	if err != nil {
		return false
	}
	opponent := b.Opponent(color)

	switch next.State {
	case opponent:
		return b.seekBottomRight(color, next)
	case color:
		if cell.State == opponent {
			return true
		}
		return false
	default:
		return false
	}
	return false
}

func (b *Board) updateBottomRight(color int, cell *Cell) {
	next, err := b.bottomRight(cell)
	if err != nil {
		return
	}
	cell.Update(color)
	opponent := b.Opponent(color)

	if next.State == opponent {
		b.updateBottomRight(color, next)
	}
}

func (b *Board) seekBottom(color int, cell *Cell) bool {
	next, err := b.bottom(cell)
	if err != nil {
		return false
	}
	opponent := b.Opponent(color)

	switch next.State {
	case opponent:
		return b.seekBottom(color, next)
	case color:
		if cell.State == opponent {
			return true
		}
		return false
	default:
		return false
	}
	return false
}

func (b *Board) updateBottom(color int, cell *Cell) {
	next, err := b.bottom(cell)
	if err != nil {
		return
	}
	cell.Update(color)
	opponent := b.Opponent(color)

	if next.State == opponent {
		b.updateBottom(color, next)
	}
}

func (b *Board) seekBottomLeft(color int, cell *Cell) bool {
	next, err := b.bottomLeft(cell)
	if err != nil {
		return false
	}
	opponent := b.Opponent(color)

	switch next.State {
	case opponent:
		return b.seekBottomLeft(color, next)
	case color:
		if cell.State == opponent {
			return true
		}
		return false
	default:
		return false
	}
	return false
}

func (b *Board) updateBottomLeft(color int, cell *Cell) {
	next, err := b.bottomLeft(cell)
	if err != nil {
		return
	}
	cell.Update(color)
	opponent := b.Opponent(color)

	if next.State == opponent {
		b.updateBottomLeft(color, next)
	}
}

func (b *Board) seekLeft(color int, cell *Cell) bool {
	next, err := b.left(cell)
	if err != nil {
		return false
	}
	opponent := b.Opponent(color)

	switch next.State {
	case opponent:
		return b.seekLeft(color, next)
	case color:
		if cell.State == opponent {
			return true
		}
		return false
	default:
		return false
	}
}

func (b *Board) updateLeft(color int, cell *Cell) {
	next, err := b.left(cell)
	if err != nil {
		return
	}
	cell.Update(color)
	opponent := b.Opponent(color)

	if next.State == opponent {
		b.updateLeft(color, next)
	}
}

func (b *Board) seekTopLeft(color int, cell *Cell) bool {
	next, err := b.topLeft(cell)
	if err != nil {
		return false
	}
	opponent := b.Opponent(color)

	switch next.State {
	case opponent:
		return b.seekTopLeft(color, next)
	case color:
		if cell.State == opponent {
			return true
		}
		return false
	default:
		return false
	}
}

func (b *Board) updateTopLeft(color int, cell *Cell) {
	next, err := b.topLeft(cell)
	if err != nil {
		return
	}
	cell.Update(color)
	opponent := b.Opponent(color)

	if next.State == opponent {
		b.updateLeft(color, next)
	}
}
