package io

import (
	"os"
	"path/filepath"
)

// AppendFile creates `path` if doesn't exist and inserts `text` into it,
// otherwise append `text` to `path`.
func AppendFile(path string, text string) (err error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	if _, err := f.Write([]byte(text)); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return
}

// NewFile creates `path` if doesn't exist and inserts `data` into it.
func NewFile(path string, data []byte) (err error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	if _, err := f.Write(data); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return
}

// FileExists check whether path file exists or not.
func FileExists(path string) (exists bool, err error) {
	info, err := os.Stat(path)
	exists = !os.IsNotExist(err) && !info.IsDir()
	if !exists {
		err = nil
	}

	return
}

// ListFiles returns a list of all files in base
func ListFiles(base string) (files []string, err error) {
	err = filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	})
	return
}
