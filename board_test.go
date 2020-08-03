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

func TestBoard_top(t *testing.T) {
	input := &Position{X: 1, Y: 2}
	expected := &Position{X: 1, Y: 1}

	b := NewBoard(InitBoard)
	cell := b.Cell(input.X, input.Y)
	actual, _ := b.top(cell)
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("got: (%d, %d)\nwant: (%d, %d)", actual.X, actual.Y, expected.X, expected.Y)
	}
}

func TestBoard_topRight(t *testing.T) {
	input := &Position{X: 1, Y: 2}
	expected := &Position{X: 2, Y: 1}

	b := NewBoard(InitBoard)
	cell := b.Cell(input.X, input.Y)
	actual, _ := b.topRight(cell)
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("got: (%d, %d)\nwant: (%d, %d)", actual.X, actual.Y, expected.X, expected.Y)
	}
}

func TestBoard_right(t *testing.T) {
	input := &Position{X: 1, Y: 2}
	expected := &Position{X: 2, Y: 2}

	b := NewBoard(InitBoard)
	cell := b.Cell(input.X, input.Y)
	actual, _ := b.right(cell)
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("got: (%d, %d)\nwant: (%d, %d)", actual.X, actual.Y, expected.X, expected.Y)
	}
}

func TestBoard_bottomRight(t *testing.T) {
	input := &Position{X: 1, Y: 2}
	expected := &Position{X: 2, Y: 3}

	b := NewBoard(InitBoard)
	cell := b.Cell(input.X, input.Y)
	actual, _ := b.bottomRight(cell)
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("got: (%d, %d)\nwant: (%d, %d)", actual.X, actual.Y, expected.X, expected.Y)
	}
}

func TestBoard_bottom(t *testing.T) {
	input := &Position{X: 1, Y: 2}
	expected := &Position{X: 1, Y: 3}

	b := NewBoard(InitBoard)
	cell := b.Cell(input.X, input.Y)
	actual, _ := b.bottom(cell)
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("got: (%d, %d)\nwant: (%d, %d)", actual.X, actual.Y, expected.X, expected.Y)
	}
}

func TestBoard_bottomLeft(t *testing.T) {
	input := &Position{X: 1, Y: 2}
	expected := &Position{X: 0, Y: 3}

	b := NewBoard(InitBoard)
	cell := b.Cell(input.X, input.Y)
	actual, _ := b.bottomLeft(cell)
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("got: (%d, %d)\nwant: (%d, %d)", actual.X, actual.Y, expected.X, expected.Y)
	}

}

func TestBoard_left(t *testing.T) {
	input := &Position{X: 1, Y: 2}
	expected := &Position{X: 0, Y: 2}

	b := NewBoard(InitBoard)
	cell := b.Cell(input.X, input.Y)
	actual, _ := b.left(cell)
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("got: (%d, %d)\nwant: (%d, %d)", actual.X, actual.Y, expected.X, expected.Y)
	}

}

func TestBoard_topLeft(t *testing.T) {
	input := &Position{X: 1, Y: 2}
	expected := &Position{X: 0, Y: 1}

	b := NewBoard(InitBoard)
	cell := b.Cell(input.X, input.Y)
	actual, _ := b.topLeft(cell)
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("got: (%d, %d)\nwant: (%d, %d)", actual.X, actual.Y, expected.X, expected.Y)
	}
}

func TestBoard_seekTop(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		expected bool
	}{
		{
			desc: "when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 3},
			expected: false,
		},
		{
			desc: "when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 3},
			expected: false,
		},
		{
			desc: "when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 3},
			expected: true,
		},
	}

	for _, tc := range testcases {
		b := NewBoard(tc.board)
		cell := b.Cell(tc.pos.X, tc.pos.Y)
		actual := b.seekTop(int(Black), cell)
		if tc.expected != actual {
			t.Errorf("%s, got: %v, expected: %v", tc.desc, actual, tc.expected)
			fmt.Println("actual")
			b.Show()
		}
	}
}

func TestBoard_seekTopRight(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		expected bool
	}{
		{
			desc: "when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 3},
			expected: false,
		},
		{
			desc: "when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 3},
			expected: false,
		},
		{
			desc: "when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 2, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 3},
			expected: true,
		},
	}

	for _, tc := range testcases {
		b := NewBoard(tc.board)
		cell := b.Cell(tc.pos.X, tc.pos.Y)
		actual := b.seekTopRight(int(Black), cell)
		if tc.expected != actual {
			t.Errorf("%s, got: %v, expected: %v", tc.desc, actual, tc.expected)
			fmt.Println("actual")
			b.Show()
		}
	}
}

func TestBoard_seekRight(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		expected bool
	}{
		{
			desc: "when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 2},
			expected: false,
		},
		{
			desc: "when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 2},
			expected: false,
		},
		{
			desc: "when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 2, 1, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 2},
			expected: true,
		},
	}

	for _, tc := range testcases {
		b := NewBoard(tc.board)
		cell := b.Cell(tc.pos.X, tc.pos.Y)
		actual := b.seekRight(int(Black), cell)
		if tc.expected != actual {
			t.Errorf("%s, got: %v, expected: %v", tc.desc, actual, tc.expected)
			fmt.Println("actual")
			b.Show()
		}
	}
}

func TestBoard_seekBottomRight(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		expected bool
	}{
		{
			desc: "when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 0},
			expected: false,
		},
		{
			desc: "when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 0},
			expected: false,
		},
		{
			desc: "when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 0, Y: 0},
			expected: true,
		},
	}

	for _, tc := range testcases {
		b := NewBoard(tc.board)
		cell := b.Cell(tc.pos.X, tc.pos.Y)
		actual := b.seekBottomRight(int(Black), cell)
		if tc.expected != actual {
			t.Errorf("%s, got: %v, expected: %v", tc.desc, actual, tc.expected)
			fmt.Println("actual")
			b.Show()
		}
	}
}

func TestBoard_seekBottom(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		expected bool
	}{
		{
			desc: "when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 0},
			expected: false,
		},
		{
			desc: "when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 0},
			expected: false,
		},
		{
			desc: "when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 2, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 1, Y: 0},
			expected: true,
		},
	}

	for _, tc := range testcases {
		b := NewBoard(tc.board)
		cell := b.Cell(tc.pos.X, tc.pos.Y)
		actual := b.seekBottom(int(Black), cell)
		if tc.expected != actual {
			t.Errorf("%s, got: %v, expected: %v", tc.desc, actual, tc.expected)
			fmt.Println("actual")
			b.Show()
		}
	}
}

func TestBoard_seekBottomLeft(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		expected bool
	}{
		{
			desc: "when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 0},
			expected: false,
		},
		{
			desc: "when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 0},
			expected: false,
		},
		{
			desc: "when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 2, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 0},
			expected: true,
		},
	}

	for _, tc := range testcases {
		b := NewBoard(tc.board)
		cell := b.Cell(tc.pos.X, tc.pos.Y)
		actual := b.seekBottomLeft(int(Black), cell)
		if tc.expected != actual {
			t.Errorf("%s, got: %v, expected: %v", tc.desc, actual, tc.expected)
			fmt.Println("actual")
			b.Show()
		}
	}
}
func TestBoard_seekLeft(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		expected bool
	}{
		{
			desc: "when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 2},
			expected: false,
		},
		{
			desc: "when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 2},
			expected: false,
		},
		{
			desc: "when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 2, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 2},
			expected: true,
		},
	}

	for _, tc := range testcases {
		b := NewBoard(tc.board)
		cell := b.Cell(tc.pos.X, tc.pos.Y)
		actual := b.seekLeft(int(Black), cell)
		if tc.expected != actual {
			t.Errorf("%s, got: %v, expected: %v", tc.desc, actual, tc.expected)
			fmt.Println("actual")
			b.Show()
		}
	}
}

func TestBoard_seekTopLeft(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		expected bool
	}{
		{
			desc: "when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 3},
			expected: false,
		},
		{
			desc: "when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 3},
			expected: false,
		},
		{
			desc: "when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 0},
			},
			pos:      &Position{X: 3, Y: 3},
			expected: true,
		},
	}

	for _, tc := range testcases {
		b := NewBoard(tc.board)
		cell := b.Cell(tc.pos.X, tc.pos.Y)
		actual := b.seekTopLeft(int(Black), cell)
		if tc.expected != actual {
			t.Errorf("%s, got: %v, expected: %v", tc.desc, actual, tc.expected)
			fmt.Println("actual")
			b.Show()
		}
	}
}

func TestBoard_updateTop(t *testing.T) {
	testcases := []struct {
		desc     string
		board    [][]int
		pos      *Position
		expected [][]int
	}{
		{
			desc: "when line end is empty cell",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			pos: &Position{X: 1, Y: 3},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
			},
		},
		{
			desc: "when line not include opponent",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			pos: &Position{X: 1, Y: 3},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 0, 0},
			},
		},
		{
			desc: "when line end is my color",
			board: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 0, 0},
			},
			pos: &Position{X: 1, Y: 3},
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
		b.updateTop(int(Black), cell)
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
