package game

import "github.com/evenboee/go-snake/utils/random"

type position struct {
	X int
	Y int
}

func newPosition(x int, y int) position {
	return position{
		X: x,
		Y: y,
	}
}

func newRandomPosition(maxX int, maxY int) position {
	return position{
		X: random.RandomInt(maxX),
		Y: random.RandomInt(maxY),
	}
}

// add adds other position to a position
//
// #virus
// Can trigger windows defender as a virus if set to public and game debug on.
// Propably because private method is executed in different package.
// Change "add" => "Add" (set to public) in case of windows defender being triggered
func (pos *position) add(other position) *position {
	pos.X += other.X
	pos.Y += other.Y
	return pos
}

// equals checks if position is equal to another
//
// see @add#virus
func (pos *position) equals(other position) bool {
	return pos.X == other.X && pos.Y == other.Y
}
