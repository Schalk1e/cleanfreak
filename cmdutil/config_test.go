package cmdutil

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v2"

	core "github.com/Schalk1e/cleanfreak/core"
)

// mockConfigInitPaths returns predefined initialization paths for testing.
//
// This avoids user interaction by simulating user input.
func mockConfigInitPaths() []string {
	return []string{"personal", "work"}
}

// mockCachePaths returns predefined cache directories for testing.
//
// Using a fixed path ensures consistent test results.
func mockCachePaths() []string {
	return []string{"~/.cache"}
}

// mockCacheSize returns a fixed cache size threshold (in GB) for testing.
//
// This replaces interactive user input with a controlled value.
func mockCacheSize() float64 {
	return 1.0
}

// TestBuildConfig verifies that BuildConfig correctly generates a YAML configuration file.
//
// It creates a temporary directory, invokes BuildConfig with mocked values,
// reads the output file, and checks if the generated configuration matches expectations.
func TestBuildConfig(t *testing.T) {
	tmpdir := t.TempDir()

	BuildConfig(
		tmpdir,
		WithConfigInitPaths(mockConfigInitPaths()),
		WithCachePaths(mockCachePaths()),
		WithCacheSize(mockCacheSize()),
	)

	configPath := filepath.Join(tmpdir, ".cleanfreak.yaml")

	data, readErr := os.ReadFile(configPath)

	var got Settings
	unmarshalErr := yaml.Unmarshal(data, &got)

	if err := errors.Join(readErr, unmarshalErr); err != nil {
		t.Fatalf("Error processing config file: %v", err)
	}

	want := Settings{
		Directory: "cleanfreak",
		Subdirs:   mockConfigInitPaths(),
		Cachedirs: mockCachePaths(),
		Threshold: mockCacheSize(),
	}

	core.ShowTestResultDeepEqual(got, want, t)
}
