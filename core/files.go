package core

import (
	"errors"
	"os"
	"path/filepath"
)

// List returns a list of all files and directories in the specified directory.
// If only_dirs is true, only directories will be included in the returned list.
// If only_dirs is false, both files and directories will be included.
func List(dir string, only_dirs bool) (files []string) {
	// Walk through the directory and its subdirectories.
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Append directories or files based on the only_dirs flag.
		if info.IsDir() {
			files = append(files, path)
		} else if !info.IsDir() && !only_dirs {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
	return
}

// IsEmpty checks whether the given directory is empty.
// It returns true if the directory contains only itself and no other files or directories.
func IsEmpty(dir string) bool {
	f_arr := List(dir, false)
	// If the length of the file array is 1, the directory is empty (contains only itself).
	return len(f_arr) == 1
}

// Clean represents a structure for cleaning files by transferring or deleting them.
type Clean struct {
	SourceFile string // The source file to be transferred or deleted.
	TargetFile string // The target file for transfer.
}

// FileTransfer moves a file from SourceFile to TargetFile.
func (clean *Clean) FileTransfer() error {
	source_file := clean.SourceFile
	target_file := clean.TargetFile

	// Attempt to rename (move) the source file to the target file path.
	err := os.Rename(source_file, target_file)
	if err != nil {
		return errors.New("failed to transfer file: " + err.Error())
	}
	return nil
}

// FileDelete deletes the SourceFile from the system.
func (clean *Clean) FileDelete() error {
	source_file := clean.SourceFile

	// Attempt to remove the file or directory.
	err := os.RemoveAll(source_file)
	if err != nil {
		return errors.New("failed to delete file: " + err.Error())
	}
	return nil
}
