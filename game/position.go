package game

import "github.com/evenboee/go-snake/utils/random"

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

func NewRandomPosition(maxX int, maxY int) Position {
	return Position{
		X: random.RandomInt(maxX),
		Y: random.RandomInt(maxY),
	}
}

func (pos *Position) Add(other Position) *Position {
	pos.X += other.X
	pos.Y += other.Y
	return pos
}

func (pos *Position) Equals(other Position) bool {
	return pos.X == other.X && pos.Y == other.Y
}
