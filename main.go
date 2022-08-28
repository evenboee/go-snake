package main

import (
	"fmt"
	"time"

	"github.com/evenboee/go-snake/game"
	"github.com/evenboee/go-snake/utils/terminal"
)

func main() {
	fmt.Println("Snake")

	terminal.Clear()

	myGame := game.NewGame()
	myGame.Run(200 * time.Millisecond)
}
