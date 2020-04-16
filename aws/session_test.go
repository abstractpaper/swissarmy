package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_, err := New("us-east-1", "test", "test")
	assert.Nil(t, err)
}
