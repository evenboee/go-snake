package utils

import (
	"log"

	"github.com/eiannone/keyboard"
)

const KeyUp = uint16(keyboard.KeyArrowUp)
const KeyRight = uint16(keyboard.KeyArrowRight)
const KeyDown = uint16(keyboard.KeyArrowDown)
const KeyLeft = uint16(keyboard.KeyArrowLeft)

type keyboardListener chan KeyEvent

type KeyEvent struct {
	Char rune
	Key  uint16
}

func NewKeyboardListener(channel chan KeyEvent) keyboardListener {
	return make(keyboardListener)
}

func (listener keyboardListener) Listen() {

	if err := keyboard.Open(); err != nil {
		log.Fatalln("Failed to open keyboard")
	}

	defer func() {
		err := keyboard.Close()
		if err != nil {
			log.Println("failed to close keyboard", err)
		} else {
			log.Println("successfully closed keyboard")
		}
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Println("failed to get key:", err)
		}

		listener <- KeyEvent{
			Char: char,
			Key:  uint16(key),
		}
	}
}

func KeyBoardListen(handleKeyEvent func(KeyEvent) bool) {
	if err := keyboard.Open(); err != nil {
		log.Fatalln("Failed to open keyboard")
	}

	defer func() {
		err := keyboard.Close()
		if err != nil {
			log.Println("failed to close keyboard", err)
		} else {
			log.Println("successfully closed keyboard")
		}
	}()

	running := true
	for running {
		log.Println("listening", running)
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Println("failed to get key:", err)
		}

		running = !handleKeyEvent(KeyEvent{
			Char: char,
			Key:  uint16(key),
		})
	}
}
