package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/evenboee/go-snake/utils"
	"github.com/evenboee/go-snake/utils/terminal"
	utilTime "github.com/evenboee/go-snake/utils/time"
)

func init() {

}

type Game struct {
	snake  snake
	food   position
	width  int
	height int
	debug  bool
}

func NewGame(debug bool) Game {
	game := Game{
		snake:  newSnake(),
		width:  9,
		height: 9,
		debug:  debug,
	}

	game.food = game.getRandomFoodSpot()

	return game
}

func (game *Game) getRandomFoodSpot() position {
	pos := newRandomPosition(game.width, game.height)

	for game.snake.head.equals(pos) || utils.Any(game.snake.tail, func(p position) bool { return p.equals(pos) }) {
		pos = newRandomPosition(game.width, game.height)
	}

	return pos
}

func (game *Game) step() bool {
	game.snake.step()
	// fmt.Println(game.snake.head)
	return game.snake.head.X >= game.width || game.snake.head.Y >= game.height
}

func (game *Game) posTo1d(pos position) int {
	return pos.Y*game.width + pos.X
}

func (game *Game) draw() {
	board := make([]uint8, game.width*game.height)

	// 0 => nothing
	// 1 => wall
	// 2 => snake head
	// 3 => snake tail
	// 4 => food

	snakeDrawTime := utilTime.Timer(func() {
		board[game.posTo1d(game.snake.head)] = 2
		for _, v := range game.snake.tail {
			board[game.posTo1d(v)] = 3
		}
	})

	board[game.posTo1d(game.food)] = 4

	var sb strings.Builder

	firstLineDrawTime := utilTime.Timer(func() {
		sb.WriteString("+")
		for i := 0; i < game.width; i++ {
			sb.WriteString("--")
		}
		sb.WriteString("+")
	})

	mainBoardDrawTime := utilTime.Timer(func() {
		for i, v := range board {
			if i%game.width == 0 {
				if i == 0 {
					sb.WriteString("\n|")
				} else {
					sb.WriteString("|\n|")
				}
			}

			tbw := "  "

			switch v {
			case 1:
				tbw = "X "
			case 2:
				tbw = "H "
			case 3:
				tbw = "T "
			case 4:
				tbw = "F "
			}

			sb.WriteString(tbw)

			// sb.WriteString(fmt.Sprintf("%d", int(i/game.width)) + fmt.Sprintf("%d", i%game.width)) // Debug. Prints cords (XY). Comment out above switch
		}
	})

	lastLineDrawTime := utilTime.Timer(func() {
		sb.WriteString("|\n")
		sb.WriteString("+")
		for i := 0; i < game.width; i++ {
			sb.WriteString("--")
		}
		sb.WriteString("+")
	})

	clearTerminalTime := utilTime.Timer(func() {
		terminal.Clear()
	})

	fmt.Println(sb.String())

	if game.debug {
		fmt.Println("Main board draw time:", mainBoardDrawTime)
		fmt.Println("Snake draw time:", snakeDrawTime)
		fmt.Println("First line draw time:", firstLineDrawTime)
		fmt.Println("Last line draw time:", lastLineDrawTime)
		fmt.Println("Clear terminal time:", clearTerminalTime)
	}
}

func (game *Game) tick(quit chan bool) {
	elapsed := utilTime.Timer(func() {
		finished := game.step()
		if !finished {
			drawTime := utilTime.Timer(func() {
				game.draw()
			})

			if game.debug {
				fmt.Println("Draw time:", drawTime)
			}
		}
		quit <- finished
	})

	if game.debug {
		fmt.Println("Time:", elapsed)
	}
}

func (game *Game) Run(speed time.Duration) {
	utilTime.SetInterval(game.tick, speed)
}
