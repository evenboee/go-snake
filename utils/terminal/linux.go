package terminal

import (
	"os"
	"os/exec"
)

type linuxTerminal struct {
	Platform string
}

func newLinuxTerminal() Terminal {
	return linuxTerminal{
		Platform: "linux",
	}
}

func (t linuxTerminal) getSize() (int, int, error) {
	return 0, 0, nil
}

func (t linuxTerminal) setSize(w int, h int) error {
	return nil
}

func (t linuxTerminal) clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
