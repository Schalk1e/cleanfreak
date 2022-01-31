package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var homedir, _ = os.UserHomeDir()
	var docs string
	var dirs []string

	docs = homedir + "/Documents"
	dirs = List(docs)
	fmt.Println(dirs)
}

func List(dir string) (files []string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return
}

func IsEmpty(dir string) bool {
	var f_arr []string

	f_arr = List(dir)
	if len(f_arr) == 1 {
		return true
	} else {
		return false
	}
}
