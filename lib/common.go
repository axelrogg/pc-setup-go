package lib

import (
	"os/exec"
)

func ExecuteCmd(name string, args ...string) (string, error) {

	out, err := exec.Command(name, args...).Output()
	return string(out), err
}
