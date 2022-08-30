package terminal

import (
	"os"
	"os/exec"
)

type windowsTerminal struct {
	Platform string
}

func newWindowsTerminal() Terminal {
	return windowsTerminal{
		Platform: "windows",
	}
}

func (t windowsTerminal) getSize() (int, int, error) {
	return 0, 0, nil
}

func (t windowsTerminal) setSize(w int, h int) error {
	return nil
}

func (t windowsTerminal) clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
