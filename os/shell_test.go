package os

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	log "github.com/sirupsen/logrus"
)

func TestExecCmd(t *testing.T) {
	type Case struct {
		Command
		ExpectedStdout []string
		ExpectedStderr []string
	}

	logger := log.New()
	cases := []Case{
		{Command{"echo -n test 1 2 3", "prfx", true, false, logger}, []string{"prfx | test 1 2 3"}, []string{""}},
		{Command{"echo -n -e test 1 2 3\ntest 4 5 6", "prfx", true, false, logger}, []string{"prfx | test 1 2 3", "prfx | test 4 5 6"}, []string{""}},
		{Command{"echo -n -e test 1 2 3\ntest 4 5 6", "", true, false, logger}, []string{"test 1 2 3", "test 4 5 6"}, []string{""}},
		{Command{"echo -n -e test 1 2 3\ntest 4 5 6", "", false, false, logger}, nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.Command.Command, func(t *testing.T) {
			stdout, _, err := tc.Execute()
			assert.Nil(t, err)
			if !reflect.DeepEqual(stdout, tc.ExpectedStdout) {
				t.Fatal("failure")
			}
		})
	}
}
