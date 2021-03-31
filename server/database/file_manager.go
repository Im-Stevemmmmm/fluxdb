package database

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var osMkdir = os.MkdirAll

// NewRelativePath constructs a relative path from a path string.
func NewRelativePath(path string) *RelativePath {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fp := fmt.Sprintf("%s/%s", dir, path)
	return (*RelativePath)(&fp)
}

// Mkdir creates a directory at the path.
func (p RelativePath) Mkdir() error {
	err := osMkdir(string(p), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// OpenFile attempts to open the file at the specified path and returns an
// error if the path is a directory and not a file or there is a problem
// opening the file.
func (p RelativePath) OpenFile() (*os.File, error) {
	path := string(p)

	fstat, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	if fstat.IsDir() {
		return nil, &OpenFileError{
			IsFile: false,
			Err:    errors.New(path + " is not a file"),
		}
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, &OpenFileError{
			IsFile: true,
			Err:    errors.New("could not open file at " + path),
		}
	}

	return f, nil
}

func (e OpenFileError) Error() string {
	return e.Err.Error()
}

// OpenFileError is an error while opening a file.
type OpenFileError struct {
	IsFile bool
	Err    error
}

// RelativePath is a path relative to the database process.
type RelativePath string
