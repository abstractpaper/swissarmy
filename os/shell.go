package os

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

// Command is a command to be executed in shell
type Command struct {
	// Command is the command to be executed
	Command string
	// Prefix adds a prefix to printed stdout/stderr
	Prefix string
	// Verbose indicates whether to print stdout/stderr
	Verbose bool
	// Logger used for logging
	Logger *log.Logger
}

// ExecCmd executes `command` and returns its output.
func ExecCmd(command *Command) (stdout []string, stderr []string, err error) {
	var wg sync.WaitGroup
	// pipes to capture stdout/err
	var stdoutReader, stdoutWriter = io.Pipe()
	var stderrReader, stderrWriter = io.Pipe()

	// split command
	args := strings.Split(command.Command, " ")

	cmd := exec.Command(args[0], args[1:]...)
	// redirect stdout and stderr to pipe
	if command.Verbose {
		cmd.Stdout = stdoutWriter
		cmd.Stderr = stderrWriter
	}

	printer := func(r io.Reader, memory *[]string, logLevel func(args ...interface{})) {
		wg.Add(1)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			var text string
			if command.Prefix != "" {
				text = fmt.Sprintf("%s | %s", command.Prefix, scanner.Text())
			} else {
				text = scanner.Text()
			}
			logLevel(text)
			*memory = append(*memory, text)
		}
		wg.Done()
	}

	if command.Verbose {
		go printer(stdoutReader, &stdout, log.Info)
		go printer(stderrReader, &stderr, log.Error)
	}

	// execute command
	if err = cmd.Start(); err != nil {
		return
	}

	if err = cmd.Wait(); err != nil {
		return
	}

	if command.Verbose {
		stdoutWriter.Close()
		stderrWriter.Close()
		wg.Wait()
	}

	return
}
