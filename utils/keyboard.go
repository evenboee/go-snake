package utils

import (
	"log"

	"github.com/eiannone/keyboard"
)

const KeyUp = uint16(keyboard.KeyArrowUp)
const KeyRight = uint16(keyboard.KeyArrowRight)
const KeyDown = uint16(keyboard.KeyArrowDown)
const KeyLeft = uint16(keyboard.KeyArrowLeft)

type KeyEvent struct {
	Char rune
	Key  uint16
}

func CloseKeyboard() {
	err := keyboard.Close()
	if err != nil {
		log.Println("failed to close keyboard", err)
	} else {
		log.Println("successfully closed keyboard")
	}
}

func KeyBoardListen(handleKeyEvent func(KeyEvent) bool) {
	if err := keyboard.Open(); err != nil {
		log.Fatalln("Failed to open keyboard")
	}

	defer CloseKeyboard()

	running := true
	for running {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Println("failed to get key:", err)
			return
		}

		running = !handleKeyEvent(KeyEvent{
			Char: char,
			Key:  uint16(key),
		})
	}
}
