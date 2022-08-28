package terminal

import (
	"os"
	"os/exec"
)

type WindowsTerminal struct {
	Platform string
}

func NewWindowsTerminal() Terminal {
	return WindowsTerminal{
		Platform: "windows",
	}
}

func (t WindowsTerminal) getSize() (int, int, error) {
	return 0, 0, nil
}

func (t WindowsTerminal) setSize(w int, h int) error {
	return nil
}

func (t WindowsTerminal) clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
