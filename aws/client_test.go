package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {
	_, err := NewSession("us-east-1", "test", "test")
	assert.Nil(t, err)
}
