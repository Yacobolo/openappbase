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

// Library represents a Context7 library to download
type Library struct {
	Name       string `json:"name"`
	Context7ID string `json:"context7_id"`
	OutputDir  string `json:"output_dir"`
	OutputFile string `json:"output_file"`
	Tokens     int    `json:"tokens"`
}

// Config represents the configuration file structure
type Config struct {
	Libraries []Library `json:"libraries"`
}

const (
	context7APIURL = "https://context7.com/api/v1"
	configPath     = "context/config.json"
	apiKeyEnv      = "CONTEXT7_API_KEY"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failure", "error", err)
		os.Exit(1)
	}
}

func run() error {
	// Load API key from environment (optional)
	apiKey := os.Getenv(apiKeyEnv)
	if apiKey == "" {
		slog.Warn("CONTEXT7_API_KEY not set - using unauthenticated API (limited documentation)")
	} else {
		slog.Info("Using authenticated Context7 API")
	}

	// Load configuration
	config, err := loadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	slog.Info("Starting Context7 documentation download", "libraries", len(config.Libraries))

	// Ensure output directories exist
	if err := createOutputDirectories(config.Libraries); err != nil {
		return fmt.Errorf("failed to create output directories: %w", err)
	}

	// Download all libraries concurrently
	if err := downloadLibraries(config.Libraries, apiKey); err != nil {
		return fmt.Errorf("failed to download libraries: %w", err)
	}

	slog.Info("Successfully downloaded all Context7 documentation")
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

func downloadLibraries(libraries []Library, apiKey string) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(libraries))

	for _, lib := range libraries {
		wg.Add(1)
		go func(l Library) {
			defer wg.Done()

			slog.Info("downloading...", "library", l.Name, "context7_id", l.Context7ID)

			if err := downloadLibrary(l, apiKey); err != nil {
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

func downloadLibrary(lib Library, apiKey string) error {
	// Build Context7 API v1 URL
	url := fmt.Sprintf("%s/%s?type=txt&tokens=%d", context7APIURL, lib.Context7ID, lib.Tokens)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Add authorization header if API key is provided
	if apiKey != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	}

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch from Context7 API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("Context7 API authentication failed - check your CONTEXT7_API_KEY")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Context7 API returned non-OK status: %s (url: %s)", resp.Status, url)
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
