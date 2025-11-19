package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// Library represents a documentation source to download
type Library struct {
	Name       string `json:"name"`
	URL        string `json:"url"`
	OutputDir  string `json:"output_dir"`
	OutputFile string `json:"output_file"`
}

// Config represents the configuration file structure
type Config struct {
	Libraries []Library `json:"libraries"`
}

const configPath = "context/config.json"

func main() {
	if err := run(); err != nil {
		slog.Error("failure", "error", err)
		os.Exit(1)
	}
}

func run() error {
	// Load configuration
	config, err := loadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	slog.Info("Starting documentation download", "libraries", len(config.Libraries))

	// Ensure output directories exist
	if err := createOutputDirectories(config.Libraries); err != nil {
		return fmt.Errorf("failed to create output directories: %w", err)
	}

	// Download all libraries concurrently
	if err := downloadLibraries(config.Libraries); err != nil {
		return fmt.Errorf("failed to download libraries: %w", err)
	}

	slog.Info("Successfully downloaded all documentation")
	return nil
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	if len(config.Libraries) == 0 {
		return nil, fmt.Errorf("no libraries defined in config")
	}

	return &config, nil
}

func createOutputDirectories(libraries []Library) error {
	dirs := make(map[string]bool)
	for _, lib := range libraries {
		dirs[lib.OutputDir] = true
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(dirs))

	for dir := range dirs {
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			if err := os.MkdirAll(d, 0755); err != nil {
				errCh <- fmt.Errorf("failed to create directory [%s]: %w", d, err)
			}
		}(dir)
	}

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func downloadLibraries(libraries []Library) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(libraries))

	for _, lib := range libraries {
		wg.Add(1)
		go func(l Library) {
			defer wg.Done()

			slog.Info("downloading...", "library", l.Name, "url", l.URL)

			if err := downloadLibrary(l); err != nil {
				errCh <- fmt.Errorf("failed to download [%s]: %w", l.Name, err)
			} else {
				slog.Info("finished", "library", l.Name, "output", filepath.Join(l.OutputDir, l.OutputFile))
			}
		}(lib)
	}

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func downloadLibrary(lib Library) error {
	// Fetch from URL
	resp, err := http.Get(lib.URL)
	if err != nil {
		return fmt.Errorf("failed to fetch from URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request returned non-OK status: %s (url: %s)", resp.Status, lib.URL)
	}

	// Prepare output file path
	outputPath := filepath.Join(lib.OutputDir, lib.OutputFile)

	// Create output file
	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file [%s]: %w", outputPath, err)
	}
	defer out.Close()

	// Copy content
	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("failed to write output file [%s]: %w", outputPath, err)
	}

	return nil
}
