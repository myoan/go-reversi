package reversi

type Cell struct {
	X     int `json:"x"`
	Y     int `json:"y"`
	State int `json:"state"`
}

type CellState int

const (
	None CellState = iota
	Black
	White
)

func (c *Cell) Update(color int) {
	c.State = color
}

func (c *Cell) noAround() bool {
	return true
}
