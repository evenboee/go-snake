package game

import (
	"fmt"
	"time"

	"github.com/evenboee/go-snake/utils/terminal"
	utilTime "github.com/evenboee/go-snake/utils/time"
)

func init() {

}

type Game struct {
	pos    Position
	dir    Position
	width  int
	height int
}

func NewGame() Game {
	return Game{
		pos:    NewPosition(0, 0),
		dir:    NewPosition(1, 1),
		width:  20,
		height: 10,
	}
}

func (game *Game) step() bool {
	game.pos.Add(game.dir)
	return game.pos.X >= game.width || game.pos.Y >= game.height
}

func (game *Game) draw() {
	terminal.Clear()

	fmt.Print("+")
	for i := 0; i < game.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	for y := 0; y < game.height; y++ {
		fmt.Print("|")

		for x := 0; x < game.width; x++ {
			if game.pos.X == x && game.pos.Y == y {
				fmt.Print("S")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println("|")
	}

	fmt.Print("+")
	for i := 0; i < game.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (game *Game) tick(quit chan bool) {
	quit <- game.step()
	game.draw()
}

func (game *Game) Run(speed time.Duration) {
	utilTime.SetInterval(game.tick, speed)
}
