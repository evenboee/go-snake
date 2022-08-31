package game

import (
	"github.com/evenboee/go-snake/utils"
)

type snake struct {
	head position
	tail []position

	dir position
}

func newSnake() snake {
	return snake{
		head: newPosition(1, 0),
		tail: []position{
			newPosition(0, 0),
		},
		dir: newPosition(1, 1),
	}
}

func (s *snake) step(remove bool) {
	if remove {
		s.tail = append(s.tail, s.head)
	} else {
		s.tail, _ = utils.ShiftIn(s.tail, s.head)
	}
	s.head.add(s.dir)
}
