package cmdutil

import (
	"fmt"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type ConfigOption func(*Settings)

// WithDefaults returns a slice of ConfigOption functions that apply default settings.
// These include default config initialization paths, cache paths, and cache size.
func WithDefaults() []ConfigOption {
	return []ConfigOption{
		WithConfigInitPaths(ConfigInitPaths()),
		WithCachePaths(CachePaths()),
		WithCacheSize(CacheSize()),
	}
}

// WithConfigInitPaths sets the initialization paths for the cleanfreak directory.
func WithConfigInitPaths(paths []string) ConfigOption {
	return func(s *Settings) {
		s.Subdirs = paths
	}
}

// WithCachePaths sets the cache directories for storing temporary data.
func WithCachePaths(paths []string) ConfigOption {
	return func(s *Settings) {
		s.Cachedirs = paths
	}
}

// WithCacheSize sets the cache size threshold in the configuration.
func WithCacheSize(size float64) ConfigOption {
	return func(s *Settings) {
		s.Threshold = size
	}
}

// Settings represents the configuration structure for the application's config.
type Settings struct {
	Directory string
	Subdirs   []string
	Cachedirs []string
	Threshold float64
}

// BuildConfig creates a YAML configuration file in the specified directory.
// It applies any provided ConfigOption functions to customize the settings.
// The configuration is written to `.cleanfreak.yaml` in the given directory.
//
// Parameters:
// - dir: The directory where the configuration file will be created.
// - opts: Variadic slice of ConfigOption functions to customize settings.
//
// Example usage:
//
//	BuildConfig("/home/user", WithCacheSize(5.0), WithConfigInitPaths([]string{"personal", "work"}))
func BuildConfig(dir string, opts ...ConfigOption) {
	fpath := path.Join(dir, ".cleanfreak.yaml")

	// Initialize default configuration
	config := Settings{
		Directory: "cleanfreak",
	}

	// Apply custom configuration options
	for _, opt := range opts {
		opt(&config)
	}

	// Convert configuration struct to YAML format
	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	// Write configuration to file
	if err := os.WriteFile(fpath, data, 0666); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Config created.")
}
