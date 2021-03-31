package database

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestRelativePath(t *testing.T) {
	// Prevent filesystem modification.
	called := false
	osMkdir = func(path string, perm os.FileMode) error {
		called = true
		return nil
	}

	rp := NewRelativePath("some/path")

	absDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Fatal("error parsing path")
	}

	path := string(*rp)
	expectedPath := fmt.Sprintf("%s/%s", absDir, "some/path")

	if path != expectedPath {
		t.Fatalf("got %s; expected %s", path, expectedPath)
	}

	err = rp.Mkdir()
	if err != nil {
		t.Fatalf("received error from new relative dir: %s", err)
	}

	if !called {
		t.Fatalf("expected called; got %v", called)
	}
}
