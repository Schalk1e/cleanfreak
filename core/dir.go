package core

import (
	"fmt"
	"os"
	"path/filepath"
)

type Dir struct {
	OS string
}

// GetDownloads returns the path to the "Downloads" directory for the current user.
// It constructs the path by combining the user's home directory with the "Downloads" subdirectory.
// If there is an error retrieving the user's home directory, the function will panic.
func (dir *Dir) GetDownloads() (downloads string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("failed to get home directory: %w", err))
	}
	downloads = filepath.Join(homedir, "Downloads")
	return
}

// GetDesktop returns the path to the "Desktop" directory for the current user.
// It constructs the path by combining the user's home directory with the "Desktop" subdirectory.
// If there is an error retrieving the user's home directory, the function will panic.
func (dir *Dir) GetDesktop() (desktop string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("failed to get home directory: %w", err))
	}
	desktop = filepath.Join(homedir, "Desktop")
	return
}

// GetTrash returns the path to the "Trash" directory based on the operating system.
// It constructs the trash path by combining the user's home directory with the appropriate subdirectory:
// - On macOS ("darwin"), it returns the `.Trash` directory.
// - On Linux, it returns the `.local/share/Trash` directory.
// If the OS is unsupported or there is an error retrieving the user's home directory, the function will panic.
func (dir *Dir) GetTrash() (trash string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("failed to get home directory: %w", err))
	}

	switch dir.OS {
	case "darwin":
		trash = filepath.Join(homedir, ".Trash")
	case "linux":
		trash = filepath.Join(homedir, ".local/share/Trash")
	default:
		panic(fmt.Errorf("unsupported OS: %s", dir.OS))
	}
	return
}

// DirExists checks if the specified directory exists on the file system.
// It returns `true` if the directory exists, and `false` if it does not.
// If an error occurs while checking, the function returns `false` without panicking.
func DirExists(dir string) (exists bool) {
	_, err := os.Stat(dir)
	if err != nil {
		exists = false
	} else {
		exists = true
	}
	return
}

// DirEmpty checks if the given directory is empty by delegating the check to the IsEmpty function.
// It returns true if the directory is empty, false otherwise.
func DirEmpty(dir string) (is_empty bool) {
	is_empty = IsEmpty(dir)
	return
}

// DirsAdd creates one or more directories, along with any necessary parent directories.
// It takes a base directory path (`baseDir`) and a slice of subdirectory names (`dirs`),
// and attempts to create each of the subdirectories inside the base directory.
// If an error occurs while creating any directory, the error is printed to the console.
func DirsAdd(baseDir string, dirs []string) {
	var fpath string

	for _, dir := range dirs {
		fpath = filepath.Join(baseDir, dir)
		err := os.MkdirAll(fpath, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// DirSize calculates the total size of the specified directory and its contents (including subdirectories).
// It returns the total size in bytes, as well as any error encountered during the traversal of the directory.
// If an error occurs while walking the directory, the error is returned alongside the calculated size (which may be zero).
func DirSize(path string) (size int64, err error) {
	err = filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return
}
