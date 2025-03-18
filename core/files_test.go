package core

import (
	"os"
	"path/filepath"
	"testing"
)

func TestList(t *testing.T) {
	// Create an empty dir to house some files.
	tmpdir, err := os.MkdirTemp("", "tmpdir")
	if err != nil {
		t.Fatal("Failed to create temp directory:", err)
	}

	// Can't have [3]string here else it conflicts with the
	// variable length type return in `got`...
	var want = []string{
		tmpdir,
		filepath.Join(tmpdir, "file1.txt"),
		filepath.Join(tmpdir, "file2.txt"),
	}

	files := map[string]string{
		"file1.txt": "Hello,",
		"file2.txt": "World!",
	}

	for name, content := range files {
		err := os.WriteFile(filepath.Join(tmpdir, name), []byte(content), 0644)
		if err != nil {
			panic(err)
		}
	}

	got := List(tmpdir, false)

	t.Cleanup(func() {
		os.RemoveAll(tmpdir)
	})

	ShowTestResultDeepEqual(got, want, t)
}
