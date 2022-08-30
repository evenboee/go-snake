package terminal

import (
	"os"
	"os/exec"
)

type windowsTerminal struct {
	Platform string
}

func newWindowsTerminal() terminal {
	return windowsTerminal{
		Platform: "windows",
	}
}

// !! NOT IMPLEMENTED
func (t windowsTerminal) getSize() (int, int, error) {
	return 0, 0, nil
}

// !! NOT IMPLEMENTED
func (t windowsTerminal) setSize(w int, h int) error {
	return nil
}

func (t windowsTerminal) clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
