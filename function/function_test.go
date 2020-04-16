package function

import (
	"errors"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestRetrySuccess(t *testing.T) {
	f := func() error {
		_ = "x"
		return nil
	}

	test := test.NewGlobal()

	Retry(f, nil)

	entries := test.AllEntries()
	assert.Equal(t, 0, len(entries))
}

func TestRetryError(t *testing.T) {
	var t1, t2 time.Time
	f := func() error {
		_ = "x"
		t1 = time.Now()

		if t2.After(t1) {
			return errors.New("test")
		}

		return nil
	}

	test := test.NewGlobal()

	t2 = time.Now().Add(5 * time.Second)
	Retry(f, nil)

	entries := test.AllEntries()
	errorCount := 0
	for _, e := range entries {
		if e.Level == logrus.ErrorLevel {
			errorCount = errorCount + 1
		}
	}
	assert.Equal(t, 2, errorCount)
}
