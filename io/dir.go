package io

import (
	"io"
	"os"
)

// DirExists checks whether path directory exists or not.
func DirExists(path string) (exists bool, err error) {
	info, err := os.Stat(path)
	exists = !os.IsNotExist(err) && info.IsDir()
	if !exists {
		err = nil
	}

	return
}

// DirEmpty checks whether path directory is empty or not.
func DirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}
