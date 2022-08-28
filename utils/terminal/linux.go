package terminal

import (
	"os"
	"os/exec"
)

type LinuxTerminal struct {
	Platform string
}

func NewLinuxTerminal() Terminal {
	return LinuxTerminal{
		Platform: "linux",
	}
}

func (t LinuxTerminal) getSize() (int, int, error) {
	return 0, 0, nil
}

func (t LinuxTerminal) setSize(w int, h int) error {
	return nil
}

func (t LinuxTerminal) clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
