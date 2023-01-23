package core

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
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
	} else if user_os == "linux" {
		trash = homedir + "/.local/share/Trash"
	}
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

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
