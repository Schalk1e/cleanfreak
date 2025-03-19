package core

import (
	"os"
	"testing"
)

// TestGetDownloads tests the GetDownloads method of the Dir struct, ensuring that it returns the correct
// path to the Downloads directory based on the user's home directory.
func TestGetDownloads(t *testing.T) {
	d := Dir{}

	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("Failed to get home directory:", err)
	}
	want := homedir + "/Downloads"

	got := d.GetDownloads()

	ShowTestResult(got, want, t)
}

// TestGetDesktop tests the GetDesktop method of the Dir struct, ensuring that it returns the correct
// path to the Desktop directory based on the user's home directory.
func TestGetDesktop(t *testing.T) {
	d := Dir{}

	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("Failed to get home directory:", err)
	}
	want := homedir + "/Desktop"

	got := d.GetDesktop()

	ShowTestResult(got, want, t)
}

// TestGetTrashDarwin tests the GetTrash method of the Dir struct for a Darwin (macOS) system,
// ensuring that it returns the correct path to the Trash directory.
func TestGetTrashDarwin(t *testing.T) {
	d := Dir{OS: "darwin"}

	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("Failed to get home directory:", err)
	}
	want := homedir + "/.Trash"

	got := d.GetTrash()

	ShowTestResult(got, want, t)
}

// TestGetTrashLinux tests the GetTrash method of the Dir struct for a Linux system,
// ensuring that it returns the correct path to the Trash directory.
func TestGetTrashLinux(t *testing.T) {
	d := Dir{OS: "linux"}

	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("Failed to get home directory:", err)
	}
	want := homedir + "/.local/share/Trash"

	got := d.GetTrash()

	ShowTestResult(got, want, t)
}

// TestDirExistsEmpty tests the DirExists function by checking if an empty temporary directory exists.
func TestDirExistsEmpty(t *testing.T) {
	// Create an empty temporary directory using t.TempDir()
	tmpdir := t.TempDir()

	got := DirExists(tmpdir)
	want := true

	ShowTestResult(got, want, t)
}

// TestDirExistsNonEmpty tests the DirExists function by checking if a non-empty temporary directory exists.
// It writes a temporary file in the directory before checking its existence.
func TestDirExistsNonEmpty(t *testing.T) {
	// Create a temporary directory using t.TempDir()
	tmpdir := t.TempDir()
	tmpfile := tmpdir + "/tmpfile.txt"

	// Write some content into file
	err := os.WriteFile(tmpfile, []byte("tmpcontent"), 0644)
	if err != nil {
		t.Fatal("Failed to write file:", err)
	}

	got := DirExists(tmpdir)
	want := true

	ShowTestResult(got, want, t)
}

// TestDirExistsNotExist tests the DirExists function by checking if a non-existing directory returns false.
func TestDirExistsNotExist(t *testing.T) {
	tmpdir := "doesnotexist"
	got := DirExists(tmpdir)
	want := false

	ShowTestResult(got, want, t)
}

// TestDirEmpty tests the DirEmpty function by checking if an empty directory is detected as empty,
// and a non-empty directory is detected as non-empty.
func TestDirEmpty(t *testing.T) {
	// Test with an empty directory
	tmpdir := t.TempDir()
	got := DirEmpty(tmpdir)
	want := true

	ShowTestResult(got, want, t)

	// Test with a non-empty directory
	tmpfile := tmpdir + "/tmpfile.txt"
	// Write some content into file
	_ = os.WriteFile(tmpfile, []byte("tmpcontent"), 0644)

	got = DirEmpty(tmpdir)
	want = false

	ShowTestResult(got, want, t)
}

// TestDirsAdd tests the DirsAdd function by creating a subdirectory within a base directory,
// and checking if it exists before and after the operation.
func TestDirsAdd(t *testing.T) {
	// Use a temporary directory as the base directory
	tmpdir := t.TempDir()
	var got [2]bool
	var want = [2]bool{false, true}
	var subdirs = []string{"subdir"}

	t.Cleanup(func() {
		os.RemoveAll(tmpdir)
	})

	got[0] = DirExists(tmpdir + "/subdir")

	DirsAdd(tmpdir, subdirs)

	got[1] = DirExists(tmpdir + "/subdir")

	ShowTestResult(got, want, t)
}

// TestDirSize tests the DirSize function by creating an empty directory and checking if its size is zero.
func TestDirSize(t *testing.T) {
	// Use a temporary directory
	tmpdir := t.TempDir()
	want := int64(0)
	var got int64

	t.Cleanup(func() {
		os.RemoveAll(tmpdir)
	})

	DirsAdd(tmpdir, []string{})

	got, _ = DirSize(tmpdir)

	ShowTestResult(got, want, t)
}
