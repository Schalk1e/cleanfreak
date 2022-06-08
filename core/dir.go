package core

import (
	"fmt"
	"os"
	"path"
)

type Dir struct {
	OS string
}

func (dir *Dir) GetDownloads() (downloads string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	downloads = homedir + "/Downloads"
	return
}

func (dir *Dir) GetDesktop() (desktop string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	desktop = homedir + "/Desktop"
	return
}

func (dir *Dir) GetTrash() (trash string) {
	user_os := dir.OS
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if user_os == "darwin" {
		trash = homedir + "/.Trash"
	} /* what happens if no list matches? */
	return
}

func DirEmpty(dir string) (is_empty bool) {
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
