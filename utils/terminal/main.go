package terminal

import (
	"runtime"
)

type terminal interface {
	// clear clears the terminal
	clear()

	// getSize returns the size of the terminal
	getSize() (int, int, error)

	// setSize sets the size of the terminal
	setSize(w int, h int) error
}

var term terminal

// initialize terminal variable based on OS
func init() {
	if runtime.GOOS == "windows" {
		term = newWindowsTerminal()
	} else {
		term = newLinuxTerminal()
	}
}

// !! NOT IMPLEMENTED
// GetSize returns the size of the console
func GetSize() (int, int, error) {
	return term.getSize()
}

// !! NOT IMPLEMENTED
// SetSize sets the size of the console
func SetSize(w int, h int) error {
	return term.setSize(w, h)
}

// Clear the console
// NOTE: Time consuming (tested to be 30ms)
func Clear() {
	term.clear()
}
