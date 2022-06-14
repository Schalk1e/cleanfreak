package core

import (
	"os"
	"path/filepath"
)

func List(dir string, only_dirs bool) (files []string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
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

func IsEmpty(dir string) bool {
	var f_arr []string

	f_arr = List(dir, false)
	if len(f_arr) == 1 {
		return true
	} else {
		return false
	}
}

type Clean struct {
	SourceFile string
	TargetFile string
}

func (clean *Clean) FileTransfer() {
	source_file := clean.SourceFile
	target_file := clean.TargetFile

	err := os.Rename(source_file, target_file)
	if err != nil {
		panic(err)
	}
}

func (clean *Clean) FileDelete() {
	source_file := clean.SourceFile

	err := os.RemoveAll(source_file)
	if err != nil {
		panic(err)
	}
}
