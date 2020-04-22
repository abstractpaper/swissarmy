package io

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	// log "github.com/sirupsen/logrus"
)

// Test the following scenarios:
// - Create + Insert
// - Append to an existing file
func TestAppendFile(t *testing.T) {
	path := "/tmp/swissarmy_test_append_file"
	text := "so much text"
	append := "even more text"

	err := os.Remove(path)

	// check if file exists
	info, err := os.Stat(path)
	exists := !os.IsNotExist(err) && !info.IsDir()
	if !exists {
		err = nil
	}
	assert.Equal(t, exists, false)

	// test create + insert
	_ = AppendFile(path, text)

	// check if file exists
	info, err = os.Stat(path)
	exists = !os.IsNotExist(err) && !info.IsDir()
	if !exists {
		err = nil
	}
	assert.Equal(t, exists, true)

	// check file contents
	b, _ := ioutil.ReadFile(path)
	assert.Equal(t, string(b), text)

	// test append
	_ = AppendFile(path, append)

	// check file contents
	b, _ = ioutil.ReadFile(path)
	assert.Equal(t, string(b), text+append)

	// clean up
	os.Remove(path)
}

// Test the following scenarios:
// - Create + Insert
// - Overwrite an existing file
func TestNewFile(t *testing.T) {
	path := "/tmp/swissarmy_test_new_file"
	text := "so much text"

	err := os.Remove(path)

	// check if file exists
	info, err := os.Stat(path)
	exists := !os.IsNotExist(err) && !info.IsDir()
	if !exists {
		err = nil
	}
	assert.Equal(t, exists, false)

	// test create + insert
	err = NewFile(path, []byte(text))
	assert.Nil(t, err)

	// check if file exists
	info, err = os.Stat(path)
	exists = !os.IsNotExist(err) && !info.IsDir()
	if !exists {
		err = nil
	}
	assert.Equal(t, exists, true)

	// check file contents
	b, _ := ioutil.ReadFile(path)
	assert.Equal(t, string(b), text)

	// test append
	err = NewFile(path, []byte(text))
	assert.Nil(t, err)

	// check file contents
	b, _ = ioutil.ReadFile(path)
	assert.Equal(t, string(b), text)

	// clean up
	os.Remove(path)
}

func TestFileExists(t *testing.T) {
	path := "/tmp/swissarmy_test_file_exists"

	// test file doesn't exist
	os.Remove(path)
	exists, _ := FileExists(path)
	assert.Equal(t, exists, false)

	// test file exists
	ioutil.WriteFile(path, []byte("random text"), 0777)
	exists, _ = FileExists(path)
	assert.Equal(t, exists, true)

	// clean up
	os.Remove(path)
}

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
