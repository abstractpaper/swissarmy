package os

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	// Stream indicates whether to stream stdout/stderr to os.Stdout
	Stream bool
	// Stdio indicates whether to use os.Stdin/out/err
	Stdio bool
	// Logger used for logging
	Logger *log.Logger
}

// Execute executes `Command` and returns its output.
func (c *Command) Execute() (stdout []string, stderr []string, err error) {
	var wg sync.WaitGroup
	// pipes to capture stdout/err
	var stdoutReader, stdoutWriter = io.Pipe()
	var stderrReader, stderrWriter = io.Pipe()

	// split command
	args := strings.Split(c.Command, " ")

	cmd := exec.Command(args[0], args[1:]...)
	// redirect stdout and stderr to pipe
	if c.Stream {
		cmd.Stdout = stdoutWriter
		cmd.Stderr = stderrWriter
	}
	// use os.Stdin/out/err
	if c.Stdio {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}

	printer := func(r io.Reader, memory *[]string, logLevel func(args ...interface{})) {
		wg.Add(1)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			var text string
			if c.Prefix != "" {
				text = fmt.Sprintf("%s | %s", c.Prefix, scanner.Text())
			} else {
				text = scanner.Text()
			}
			logLevel(text)
			*memory = append(*memory, text)
		}
		wg.Done()
	}

	if c.Stream {
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

	if c.Stream {
		stdoutWriter.Close()
		stderrWriter.Close()
		wg.Wait()
	}

	return
}
