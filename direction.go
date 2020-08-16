package reversi

type Direction struct {
	dx int
	dy int
}

func (d *Direction) Next(x, y int) (int, int) {
	dx := x + d.dx
	dy := y + d.dy
	return dx, dy
}
