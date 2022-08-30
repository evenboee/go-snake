package terminal

import (
	"runtime"
)

type Terminal interface {
	clear()
	getSize() (int, int, error)
	setSize(w int, h int) error
}

var terminal Terminal

func init() {

	if runtime.GOOS == "windows" {
		terminal = newWindowsTerminal()
	} else {
		terminal = newLinuxTerminal()
	}
}

func GetSize() (int, int, error) {
	return terminal.getSize()
}

func SetSize(w int, h int) error {
	return terminal.setSize(w, h)
}

func Clear() {
	terminal.clear()
}
