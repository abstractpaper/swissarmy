package os

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecCmd(t *testing.T) {
	cmd := `echo`
	args := []string{"-n", `"test 1 2 3"`}
	out, _ := ExecCmd(cmd, args...)
	assert.Equal(t, `"test 1 2 3"`, out)
}
