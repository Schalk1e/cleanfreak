package core

import (
	"fmt"
	"os"
	"path"
)

func Dir(dir_type string) (dir string) {
	if dir_type != "Downloads" && dir_type != "Documents" && dir_type != "Desktop" {
		fmt.Println("Error: dir_type must be one of 'Downloads' or 'Documents' or 'Desktop'.")
		os.Exit(1)
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

func DirAdd(base_dir string, dirs []string) {
	cf_root := "cleanfreak" // Get this from config...
	var fpath string

	if DirExists(cf_root) {
		response := fmt.Sprintf("Already contains a directory named %s!", cf_root)
		fmt.Println(response)
	} else {
		for _, dir := range dirs {
			fpath = path.Join(base_dir, cf_root, dir)
			err := os.MkdirAll(fpath, 0755)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
