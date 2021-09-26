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
