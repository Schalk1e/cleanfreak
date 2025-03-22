package cmdutil

import (
	"fmt"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

// NOTE: Defines a type which is a function that takes
// a pointer to settings and returns nothing.
type ConfigOption func(*Settings)

func WithDefaults() []ConfigOption {
	return []ConfigOption{
		WithConfigInitPaths(ConfigInitPaths()),
		WithCachePaths(CachePaths()),
		WithCacheSize(CacheSize()),
	}
}

func WithConfigInitPaths(paths []string) ConfigOption {
	return func(s *Settings) {
		s.Subdirs = paths
	}
}

func WithCachePaths(paths []string) ConfigOption {
	return func(s *Settings) {
		s.Cachedirs = paths
	}
}

func WithCacheSize(size float64) ConfigOption {
	return func(s *Settings) {
		s.Threshold = size
	}
}

type Settings struct {
	Directory string
	Subdirs   []string
	Cachedirs []string
	Threshold float64
}

func BuildConfig(dir string, opts ...ConfigOption) {
	fpath := path.Join(dir, ".cleanfreak.yaml")

	config := Settings{
		Directory: "cleanfreak",
	}

	for _, opt := range opts {
		opt(&config)
	}

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
