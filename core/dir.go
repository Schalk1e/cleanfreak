package core

import (
	"fmt"
	"os"
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
