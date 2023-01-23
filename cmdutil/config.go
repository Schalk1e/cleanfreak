package cmdutil

import (
	"fmt"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	Directory string
	Subdirs   []string
	Cachedirs []string
}

func BuildConfig(dir string) {
	fpath := path.Join(dir, ".cleanfreak.yaml")

	// Ask user for initialisation paths.
	config_init_paths := ConfigInitPaths()
	// Ask user for cache folders to monitor.
	cache_paths := CachePaths()

	config := Settings{Directory: "cleanfreak", Subdirs: config_init_paths, Cachedirs: cache_paths}
	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	err2 := os.WriteFile(fpath, data, 0666)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Config created.")
}
