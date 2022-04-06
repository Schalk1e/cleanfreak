package core

import (
	"fmt"
	"os"
	"path"
)

func Dir(dir_type string) (dir string) {
	if dir_type != "Downloads" && dir_type != "Documents" && dir_type != "Desktop" {
		panic("Error: dir_type must be one of 'Downloads' or 'Documents' or 'Desktop'.")
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dir = homedir + "/" + dir_type
	return
}

func DirEmpty(dir_type string) (is_empty bool) {
	var dir string

	dir = Dir(dir_type)
	is_empty = IsEmpty(dir)

	return
}

func DirExists(dir string) (exists bool) {
	_, err := os.Stat(dir)
	if err != nil {
		exists = false
	} else {
		exists = true
	}
	return
}

func DirsAdd(base_dir string, dirs []string) {
	var fpath string

	for _, dir := range dirs {
		fpath = path.Join(base_dir, dir)
		err := os.MkdirAll(fpath, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}

}
