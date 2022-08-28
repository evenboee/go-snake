package game

type Position struct {
	X int
	Y int
}

func NewPosition(x int, y int) Position {
	return Position{
		X: x,
		Y: y,
	}
}

func (pos *Position) Add(other Position) *Position {
	pos.X += other.X
	pos.Y += other.Y
	return pos
}
