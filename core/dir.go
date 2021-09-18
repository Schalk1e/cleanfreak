package core

import (
	"fmt"
	"os"
)

func Dir(dir_type string) (dir string) {
	if dir_type != "downloads" && dir_type != "documents" {
		fmt.Println("Error: dir_type must be one of 'downloads' or 'documents'.")
		os.Exit(1)
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if dir_type == "downloads" {
		dir = homedir + "/Downloads"
	} else if dir_type == "documents" {
		dir = homedir + "/Documents"
	}
	return
}
