package main

import (
	"time"

	"github.com/evenboee/go-snake/game"
	"github.com/evenboee/go-snake/utils/terminal"
)

func main() {
	terminal.Clear()

	myGame := game.NewGame(true)
	myGame.Run(200 * time.Millisecond)
}
