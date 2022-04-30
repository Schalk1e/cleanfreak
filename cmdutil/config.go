package cmdutil

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	Directory string
	Subdirs   [2]string
}

func BuildConfig(dir string) {
	var fpath string

	fpath = path.Join(dir, ".cleanfreak.yaml")

	config := Settings{Directory: "cleanfreak", Subdirs: [2]string{"Personal", "Work"}}
	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	err2 := ioutil.WriteFile(fpath, data, 0666)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Config created.")
}
