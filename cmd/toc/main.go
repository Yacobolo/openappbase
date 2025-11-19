package main

// TODO: change output file name to summary.md

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	maxDepth := flag.Int("depth", 3, "Maximum header depth to include (1-6)")
	flag.Parse()

	if *maxDepth < 1 || *maxDepth > 6 {
		slog.Error("Invalid depth", "depth", *maxDepth)
		os.Exit(1)
	}

	contextDir := "context"

	// Get all subdirectories in context/
	entries, err := os.ReadDir(contextDir)
	if err != nil {
		slog.Error("Failed to read context directory", "error", err)
		os.Exit(1)
	}

	processedCount := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		subDir := filepath.Join(contextDir, entry.Name())
		mdFiles, err := filepath.Glob(filepath.Join(subDir, "*.md"))
		if err != nil {
			slog.Error("Failed to glob markdown files", "dir", subDir, "error", err)
			continue
		}

		for _, mdFile := range mdFiles {
			// Skip toc.md files
			if filepath.Base(mdFile) == "toc.md" {
				continue
			}

			if err := generateTOC(mdFile, *maxDepth); err != nil {
				slog.Error("Failed to generate TOC", "file", mdFile, "error", err)
				continue
			}
			processedCount++
			slog.Info("Generated TOC", "file", mdFile, "depth", *maxDepth)
		}
	}

	if processedCount == 0 {
		slog.Warn("No markdown files found to process")
	} else {
		slog.Info("TOC generation complete", "files_processed", processedCount)
	}
}

func generateTOC(mdFilePath string, maxDepth int) error {
	// Read the markdown file to extract headers
	file, err := os.Open(mdFilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var headers []struct {
		lineNum int
		level   int
		text    string
	}

	scanner := bufio.NewScanner(file)
	lineNum := 0
	totalLines := 0
	for scanner.Scan() {
		lineNum++
		totalLines++
		line := scanner.Text()

		// Check if line starts with #
		if strings.HasPrefix(line, "#") {
			// Count the number of # characters (header level)
			level := len(line) - len(strings.TrimLeft(line, "#"))

			// Skip if deeper than max depth
			if level > maxDepth {
				continue
			}

			// Extract the header text (remove leading # and whitespace)
			text := strings.TrimSpace(strings.TrimLeft(line, "#"))
			headers = append(headers, struct {
				lineNum int
				level   int
				text    string
			}{lineNum, level, text})
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan file: %w", err)
	}

	// Generate TOC file
	tocFilePath := filepath.Join(filepath.Dir(mdFilePath), "toc.md")
	tocFile, err := os.Create(tocFilePath)
	if err != nil {
		return fmt.Errorf("failed to create TOC file: %w", err)
	}
	defer tocFile.Close()

	writer := bufio.NewWriter(tocFile)

	// Write header
	baseName := filepath.Base(mdFilePath)
	fmt.Fprintf(writer, "# File Map: %s\n\n", baseName)

	// Write TOC entries with hierarchy and line counts
	for i, h := range headers {
		// Indent based on level (level 1 = no indent, level 2 = 2 spaces, etc.)
		indent := strings.Repeat("  ", h.level-1)

		// Calculate line count to next section at same or higher level
		lineCount := 0
		for j := i + 1; j < len(headers); j++ {
			if headers[j].level <= h.level {
				lineCount = headers[j].lineNum - h.lineNum
				break
			}
		}
		// If no next section found, calculate to end of file
		if lineCount == 0 {
			lineCount = totalLines - h.lineNum + 1
		}

		// Format: - [295:413] Header Text (start:count)
		fmt.Fprintf(writer, "%s- [%d:%d] %s\n", indent, h.lineNum, lineCount, h.text)
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to write TOC file: %w", err)
	}

	return nil
}
