package core

import (
	"os"
	"path/filepath"
	"testing"
)

// TestList tests the List function by creating temporary files in a directory,
// calling the List function to retrieve the directory contents, and comparing
// the result to the expected output.
func TestList(t *testing.T) {
	tmpdir := t.TempDir()
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

	// Create files in the temporary directory
	for name, content := range files {
		err := os.WriteFile(filepath.Join(tmpdir, name), []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to write file: %v", err)
		}
	}

	got := List(tmpdir, false)

	// Show the result
	ShowTestResultDeepEqual(got, want, t)
}

// TestIsEmptyDirEmpty tests the IsEmpty function on an empty directory,
// expecting the result to be true.
func TestIsEmptyDirEmpty(t *testing.T) {
	tmpdir := t.TempDir()

	want := true
	got := IsEmpty(tmpdir)

	ShowTestResult(got, want, t)
}

// TestIsEmptyDirNotEmpty tests the IsEmpty function on a directory
// that contains a file, expecting the result to be false.
func TestIsEmptyDirNotEmpty(t *testing.T) {
	tmpdir := t.TempDir()

	// Create a file in the directory to make it non-empty
	err := os.WriteFile(filepath.Join(tmpdir, "file.txt"), []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	want := false
	got := IsEmpty(tmpdir)

	ShowTestResult(got, want, t)
}

// TestFileTransfer tests the FileTransfer method by moving a file from
// one location to another within the same directory and checking that the
// transfer was successful.
func TestFileTransfer(t *testing.T) {
	tmpdir := t.TempDir()

	// Create a file in the directory
	err := os.WriteFile(filepath.Join(tmpdir, "file.txt"), []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	// Transfer file (rename operation)
	c := Clean{
		SourceFile: filepath.Join(tmpdir, "file.txt"),
		TargetFile: filepath.Join(tmpdir, "new_file.txt"),
	}

	want := []string{tmpdir, filepath.Join(tmpdir, "new_file.txt")}

	c.FileTransfer()

	got := List(tmpdir, false)

	ShowTestResultDeepEqual(got, want, t)
}

// TestFileDelete tests the FileDelete method by deleting a file from a directory
// and checking that the file is successfully removed.
func TestFileDelete(t *testing.T) {
	tmpdir := t.TempDir()

	// Create a file in the directory
	err := os.WriteFile(filepath.Join(tmpdir, "file.txt"), []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	c := Clean{
		SourceFile: filepath.Join(tmpdir, "file.txt"),
	}

	want := []string{tmpdir}

	c.FileDelete()

	got := List(tmpdir, false)

	ShowTestResultDeepEqual(got, want, t)
}
