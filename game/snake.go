package game

import (
	"github.com/evenboee/go-snake/utils"
)

type snake struct {
	head Position
	tail []Position

	dir Position
}

func newSnake() snake {
	return snake{
		head: NewPosition(1, 0),
		tail: []Position{
			NewPosition(0, 0),
		},
		dir: NewPosition(1, 1),
	}
}

func (s *snake) step() {
	s.tail, _ = utils.ShiftIn(s.tail, s.head)
	s.head.Add(s.dir)
}
