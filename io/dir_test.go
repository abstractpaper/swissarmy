package io

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	// log "github.com/sirupsen/logrus"
)

func TestDirExists(t *testing.T) {
	path := "/tmp/swissarmy_test_dir_exists/"

	// test dir doesn't exist
	os.RemoveAll(path)
	exists, _ := DirExists(path)
	assert.Equal(t, exists, false)

	// test file exists
	os.Mkdir(path, 0600)
	exists, _ = DirExists(path)
	assert.Equal(t, exists, true)

	// clean up
	os.RemoveAll(path)
}

func TestDirEmpty(t *testing.T) {
	path := "/tmp/swissarmy_test_dir_empty"
	file := path + "/test.txt"

	// test not empty
	os.RemoveAll(path)
	os.Mkdir(path, 0777)
	ioutil.WriteFile(file, []byte("random text"), 0777)
	empty, _ := DirEmpty(path)
	assert.Equal(t, empty, false)

	// test empty
	os.RemoveAll(path)
	os.Mkdir(path, 0600)
	empty, _ = DirEmpty(path)
	assert.Equal(t, empty, true)

	// clean up
	os.RemoveAll(path)
}
