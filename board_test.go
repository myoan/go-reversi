package reversi

import (
	"fmt"
	"testing"
)

func matchArray(src, dst [][]int) bool {
	if len(src) != len(dst) {
		return false
	}
	for i, srcLine := range src {
		if len(srcLine) != len(dst[i]) {
			return false
		}

		for j, elem := range srcLine {
			if elem != dst[i][j] {
				return false
			}
		}
	}
	return true
}

func TestBoard_Cell(t *testing.T) {
	testcases := []struct {
		desc     string
		x        int
		y        int
		expected *Cell
	}{
		{
			desc:     "when valid index",
			x:        1,
			y:        1,
			expected: &Cell{X: 1, Y: 1, State: 5},
		},
		{
			desc:     "when index is negative",
			x:        -1,
			y:        1,
			expected: nil,
		},
		{
			desc:     "when index is negative",
			x:        1,
			y:        -1,
			expected: nil,
		},
		{
			desc:     "when over width",
			x:        9,
			y:        1,
			expected: nil,
		},
		{
			desc:     "when over height",
			x:        1,
			y:        9,
			expected: nil,
		},
	}
	b := NewBoard([][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 14, 14},
	})
	for _, tc := range testcases {
		actual := b.Cell(tc.x, tc.y)
		if tc.expected == nil {
			if actual != nil {
				t.Errorf("%s, got: %v, expected: %v", tc.desc, actual, tc.expected)
			}
		} else if tc.expected.State != actual.State {
			t.Errorf("%s, got: %v, expected: %v", tc.desc, actual, tc.expected)
		}
	}
}

func TestBoard_next(t *testing.T) {
	testcases := []struct {
		desc     string
		input    *Position
		d        *Direction
		expected *Position
	}{
		{
			desc:     "direction: top",
			input:    &Position{X: 1, Y: 2},
			d:        &Direction{dx: 0, dy: -1},
			expected: &Position{X: 1, Y: 1},
		},
		{
			desc:     "direction: top-right",
			input:    &Position{X: 1, Y: 2},
			d:        &Direction{dx: 1, dy: -1},
			expected: &Position{X: 2, Y: 1},
		},
		{
			desc:     "direction: right",
			input:    &Position{X: 1, Y: 2},
			d:        &Direction{dx: 1, dy: 0},
			expected: &Position{X: 2, Y: 2},
		},
		{
			desc:     "direction: bottom right",
			input:    &Position{X: 1, Y: 2},
			d:        &Direction{dx: 1, dy: 1},
			expected: &Position{X: 2, Y: 3},
		},
		{
			desc:     "direction: bottom right",
			input:    &Position{X: 1, Y: 2},
			d:        &Direction{dx: 1, dy: 1},
			expected: &Position{X: 2, Y: 3},
		},
		{
			desc:     "direction: bottom",
			input:    &Position{X: 1, Y: 2},
			d:        &Direction{dx: 0, dy: 1},
			expected: &Position{X: 1, Y: 3},
		},
		{
			desc:     "direction: bottom left",
			input:    &Position{X: 1, Y: 2},
			d:        &Direction{dx: -1, dy: 1},
			expected: &Position{X: 0, Y: 3},
		},
		{
			desc:     "direction: left",
			input:    &Position{X: 1, Y: 2},
			d:        &Direction{dx: -1, dy: 0},
			expected: &Position{X: 0, Y: 2},
		},
		{
			desc:     "direction: top left",
			input:    &Position{X: 1, Y: 2},
			d:        &Direction{dx: -1, dy: -1},
			expected: &Position{X: 0, Y: 1},
		},
	}

	b := NewBoard(InitBoard)
	for _, tc := range testcases {
		cell := b.Cell(tc.input.X, tc.input.Y)
		actual, _ := b.next(tc.d, cell)
		if actual.X != tc.expected.X || actual.Y != tc.expected.Y {
			t.Errorf("got: (%d, %d)\nwant: (%d, %d)", actual.X, actual.Y, tc.expected.X, tc.expected.Y)
		}
	}
}
func TestBoard_seek(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		d        *Direction
		expected bool
	}{
		{
			desc: "d: top, when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 3},
			d:        &Direction{dx: 0, dy: -1},
			expected: false,
		},
		{
			desc: "d: top, when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 3},
			d:        &Direction{dx: 0, dy: -1},
			expected: false,
		},
		{
			desc: "d: top, when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 3},
			d:        &Direction{dx: 0, dy: -1},
			expected: true,
		},
		{
			desc: "d: top-right, when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 3},
			d:        &Direction{dx: 1, dy: -1},
			expected: false,
		},
		{
			desc: "d: top-right, when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 3},
			d:        &Direction{dx: 1, dy: -1},
			expected: false,
		},
		{
			desc: "d: top-right, when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 2, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 3},
			d:        &Direction{dx: 1, dy: -1},
			expected: true,
		},
		{
			desc: "d: left, when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 2},
			d:        &Direction{dx: 1, dy: 0},
			expected: false,
		},
		{
			desc: "d: left, when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 2},
			d:        &Direction{dx: 1, dy: 0},
			expected: false,
		},
		{
			desc: "d: left, when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 2, 1, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 2},
			d:        &Direction{dx: 1, dy: 0},
			expected: true,
		},
		{
			desc: "d: bottom-right, when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 0},
			d:        &Direction{dx: 1, dy: 1},
			expected: false,
		},
		{
			desc: "d: bottom-right, when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 0},
			d:        &Direction{dx: 1, dy: 1},
			expected: false,
		},
		{
			desc: "d: bottom-right, when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 0},
			d:        &Direction{dx: 1, dy: 1},
			expected: true,
		},
		{
			desc: "d: bottom, when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 0},
			d:        &Direction{dx: 0, dy: 1},
			expected: false,
		},
		{
			desc: "d: bottom, when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 0},
			d:        &Direction{dx: 0, dy: 1},
			expected: false,
		},
		{
			desc: "d: bottom, when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 2, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 0},
			d:        &Direction{dx: 0, dy: 1},
			expected: true,
		},
		{
			desc: "d: bottom-left, when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 0},
			d:        &Direction{dx: -1, dy: 1},
			expected: false,
		},
		{
			desc: "d: bottom-left, when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 0},
			d:        &Direction{dx: -1, dy: 1},
			expected: false,
		},
		{
			desc: "d: bottom-left, when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 2, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 0},
			d:        &Direction{dx: -1, dy: 1},
			expected: true,
		},
		{
			desc: "d: left, when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 2},
			d:        &Direction{dx: -1, dy: 0},
			expected: false,
		},
		{
			desc: "d: left, when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 2},
			d:        &Direction{dx: -1, dy: 0},
			expected: false,
		},
		{
			desc: "d: left, when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 2, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 2},
			d:        &Direction{dx: -1, dy: 0},
			expected: true,
		},
		{
			desc: "d: top-left, when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 3},
			d:        &Direction{dx: -1, dy: -1},
			expected: false,
		},
		{
			desc: "d: top-left, when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 3},
			d:        &Direction{dx: -1, dy: -1},
			expected: false,
		},
		{
			desc: "d: top-left, when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 3},
			d:        &Direction{dx: -1, dy: -1},
			expected: true,
		},
	}

	for _, tc := range testcases {
		b := NewBoard(tc.board)
		cell := b.Cell(tc.pos.X, tc.pos.Y)
		actual := b.seek(tc.d, int(Black), cell)
		if tc.expected != actual {
			t.Errorf("%s, got: %v, expected: %v", tc.desc, actual, tc.expected)
			fmt.Println("actual")
			b.Show()
		}
	}
}
func TestBoard_update(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		d        *Direction
		expected [][]int
	}{
		{
			desc: "d: top, when line end is board end",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos: &Position{X: 1, Y: 0},
			d:   &Direction{dx: 0, dy: -1},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			desc: "d: top, when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos: &Position{X: 1, Y: 3},
			d:   &Direction{dx: 0, dy: -1},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
			},
		},
		{
			desc: "d: top, when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos: &Position{X: 1, Y: 3},
			d:   &Direction{dx: 0, dy: -1},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 0, 0},
			},
		},
		{
			desc: "d: top, when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 0, 0},
			},
			pos: &Position{X: 1, Y: 3},
			d:   &Direction{dx: 0, dy: -1},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 0, 0},
			},
		},
	}
	for _, tc := range testcases {
		b := NewBoard(tc.board)
		cell := b.Cell(tc.pos.X, tc.pos.Y)
		b.update(tc.d, int(Black), cell)
		if !matchArray(tc.expected, b.toArray()) {
			t.Errorf("%s", tc.desc)
			fmt.Println("actual")
			b.Show()
		}
	}
}

func TestBoard_allocate(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		expected [][]int
	}{
		{
			desc: "when simple test",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 2, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos: &Position{X: 1, Y: 0},
			expected: [][]int{
				{0, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			desc: "when dual line",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 2, 2, 0},
				{0, 1, 0, 1},
				{0, 0, 0, 0},
			},
			pos: &Position{X: 1, Y: 0},
			expected: [][]int{
				{0, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 0, 1},
				{0, 0, 0, 0},
			},
		},
		{
			desc: "when some line finish no my color",
			board: [][]int{
				{0, 0, 2, 2},
				{0, 2, 2, 0},
				{0, 1, 0, 1},
				{0, 0, 0, 0},
			},
			pos: &Position{X: 1, Y: 0},
			expected: [][]int{
				{0, 1, 2, 2},
				{0, 1, 1, 0},
				{0, 1, 0, 1},
				{0, 0, 0, 0},
			},
		},
	}

	for _, tc := range testcases {
		b := NewBoard(tc.board)
		cell := b.Cell(tc.pos.X, tc.pos.Y)
		b.allocate(int(Black), cell)
		if !matchArray(tc.expected, b.toArray()) {
			t.Errorf("%s", tc.desc)
			fmt.Println("actual")
			b.Show()
		}
	}
}

func TestBoard_Count(t *testing.T) {
	testcases := []struct {
		desc   string
		board  [][]int
		expect int
	}{
		{
			desc: "empty board",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			expect: 0,
		},
		{
			desc: "when simple test",
			board: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			expect: 1,
		},
		{
			desc: "when exists multi color",
			board: [][]int{
				{1, 0, 0, 0},
				{0, 0, 1, 0},
				{2, 2, 0, 0},
				{0, 0, 0, 0},
			},
			expect: 2,
		},
	}

	for _, tc := range testcases {
		b := NewBoard(tc.board)
		actual := b.Count(int(Black))
		if actual != tc.expect {
			t.Errorf("%s", tc.desc)
			fmt.Printf("expect: %d, actual: %d\n", tc.expect, actual)
		}
	}
}
