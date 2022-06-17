package lib

import (
	"os/exec"
)

func ExecuteCmd(name string, args ...string) (string, error) {

	cmd := exec.Command(name, args...)

	out, err := cmd.CombinedOutput()

	return string(out), err
}
