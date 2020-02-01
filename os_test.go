package swissarmy

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	key := "SWISSARMY_TEST_ENV_VAR"

	// check for unset env variable
	os.Unsetenv(key)
	exists := GetEnv(key, "default value")
	assert.Equal(t, exists, "default value")

	// check for set env variable
	os.Setenv(key, "abc")
	exists = GetEnv(key, "default value")
	assert.Equal(t, exists, "abc")

	// clean up
	os.Unsetenv(key)
}
