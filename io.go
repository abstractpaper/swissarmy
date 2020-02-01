package swissarmy

import (
	"io"
	"os"
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

// FileExists check whether path file exists or not.
func FileExists(path string) (exists bool, err error) {
	info, err := os.Stat(path)
	exists = !os.IsNotExist(err) && !info.IsDir()
	if !exists {
		err = nil
	}

	return
}

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
