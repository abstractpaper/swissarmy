package io

import (
	"fmt"
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
	assert.Equal(t, false, exists)

	// test create + insert
	_ = AppendFile(path, text)

	// check if file exists
	info, err = os.Stat(path)
	exists = !os.IsNotExist(err) && !info.IsDir()
	if !exists {
		err = nil
	}
	assert.Equal(t, true, exists)

	// check file contents
	b, _ := ioutil.ReadFile(path)
	assert.Equal(t, text, string(b))

	// test append
	_ = AppendFile(path, append)

	// check file contents
	b, _ = ioutil.ReadFile(path)
	assert.Equal(t, text+append, string(b))

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
	assert.Equal(t, false, exists)

	// test create + insert
	err = NewFile(path, []byte(text))
	assert.Nil(t, err)

	// check if file exists
	info, err = os.Stat(path)
	exists = !os.IsNotExist(err) && !info.IsDir()
	if !exists {
		err = nil
	}
	assert.Equal(t, true, exists)

	// check file contents
	b, _ := ioutil.ReadFile(path)
	assert.Equal(t, text, string(b))

	// test append
	err = NewFile(path, []byte(text))
	assert.Nil(t, err)

	// check file contents
	b, _ = ioutil.ReadFile(path)
	assert.Equal(t, text, string(b))

	// clean up
	os.Remove(path)
}

func TestFileExists(t *testing.T) {
	path := "/tmp/swissarmy_test_file_exists"

	// test file doesn't exist
	os.Remove(path)
	exists, _ := FileExists(path)
	assert.Equal(t, false, exists)

	// test file exists
	ioutil.WriteFile(path, []byte("random text"), 0777)
	exists, _ = FileExists(path)
	assert.Equal(t, true, exists)

	// clean up
	os.Remove(path)
}

func TestListFiles(t *testing.T) {
	base := "/tmp/swissarmy_test_list_files/"
	path1 := fmt.Sprintf("%s/file_1", base)
	path2 := fmt.Sprintf("%s/file_2", base)

	// create folder and files
	os.Mkdir(base, 0777)
	ioutil.WriteFile(path1, []byte("random text"), 0777)
	ioutil.WriteFile(path2, []byte("random text"), 0777)

	// list files
	files, _ := ListFiles(base)
	assert.Equal(t, 2, len(files))

	// clean up
	os.RemoveAll(base)
}

func TestTopLevelFiles(t *testing.T) {
	base := "/tmp/swissarmy_test_top_level_files/"
	folder := fmt.Sprintf("%s/folder1/", base)
	path1 := fmt.Sprintf("%s/file_1", base)
	path2 := fmt.Sprintf("%s/file_2", folder)
	path3 := fmt.Sprintf("%s/file_3", folder)

	// create folder and files
	os.Mkdir(base, 0777)
	os.Mkdir(folder, 0777)
	ioutil.WriteFile(path1, []byte("random text"), 0777)
	ioutil.WriteFile(path2, []byte("random text"), 0777)
	ioutil.WriteFile(path3, []byte("random text"), 0777)

	// list files
	files, _ := TopLevelFiles(base)
	assert.Equal(t, 1, len(files))

	// clean up
	os.RemoveAll(base)
}
