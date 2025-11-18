package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func main() {
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

			if err := generateTOC(mdFile); err != nil {
				slog.Error("Failed to generate TOC", "file", mdFile, "error", err)
				continue
			}
			processedCount++
			slog.Info("Generated TOC", "file", mdFile)
		}
	}

	if processedCount == 0 {
		slog.Warn("No markdown files found to process")
	} else {
		slog.Info("TOC generation complete", "files_processed", processedCount)
	}
}

func generateTOC(mdFilePath string) error {
	// Read the markdown file to count total lines and extract headers
	file, err := os.Open(mdFilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var headers []struct {
		lineNum int
		text    string
	}

	scanner := bufio.NewScanner(file)
	lineNum := 0
	totalLines := 0
	for scanner.Scan() {
		lineNum++
		totalLines = lineNum
		line := scanner.Text()

		// Check if line starts with #
		if strings.HasPrefix(line, "#") {
			// Extract the header text (remove leading # and whitespace)
			text := strings.TrimSpace(strings.TrimLeft(line, "#"))
			headers = append(headers, struct {
				lineNum int
				text    string
			}{lineNum, text})
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
	fmt.Fprintf(writer, "# Table of Contents for %s\n\n", baseName)
	fmt.Fprintf(writer, "This file is auto-generated. Run `task context:toc` to regenerate.\n\n")

	// Write TOC entries with line ranges
	fmt.Fprintln(writer, "## Contents\n")
	for i, h := range headers {
		startLine := h.lineNum
		var endLine int

		// Calculate end line based on next header or EOF
		if i < len(headers)-1 {
			endLine = headers[i+1].lineNum - 1
		} else {
			endLine = totalLines
		}

		// Format: L0001-L0017 | Header Text
		fmt.Fprintf(writer, "L%04d-L%04d | %s\n", startLine, endLine, h.text)
	}

	// Write usage examples
	fmt.Fprintln(writer, "\n## Usage Examples\n")
	fmt.Fprintln(writer, "```bash")
	fmt.Fprintf(writer, "# Read a specific section by line range\n")
	fmt.Fprintf(writer, "sed -n '54,102p' %s\n\n", mdFilePath)
	fmt.Fprintf(writer, "# Search for specific topics\n")
	fmt.Fprintf(writer, "grep -n -i 'install' %s\n\n", mdFilePath)
	fmt.Fprintf(writer, "# Show context around a match (5 lines before and after)\n")
	fmt.Fprintf(writer, "grep -n -C 5 'example' %s\n", mdFilePath)
	fmt.Fprintln(writer, "```")

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to write TOC file: %w", err)
	}

	return nil
}
