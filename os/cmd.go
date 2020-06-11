package os

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// ExecCmd executes `command` with `args` and returns its output.
func ExecCmd(command string, args ...string) (out string, err error) {
	// pipe to capture stdout/err
	r, w, _ := os.Pipe()

	// execute cmd with args
	cmd := exec.Command(command, args...)
	// redirect stdout and stderr to pipe
	cmd.Stdout = w
	cmd.Stderr = w

	err = cmd.Run()
	if err != nil {
		return
	}
	w.Close()

	// read from pipe
	bytes, _ := ioutil.ReadAll(r)
	out = string(bytes)

	return
}
