package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

// Package represents a Go package to document
type Package struct {
	ImportPath string // Full import path (e.g., "github.com/go-chi/chi/v5")
	OutputName string // Output filename without extension (e.g., "chi")
}

func main() {
	// Define packages to document
	packages := []Package{
		{
			ImportPath: "github.com/starfederation/datastar-go/datastar",
			OutputName: "datastar-go",
		},
		{
			ImportPath: "github.com/go-chi/chi/v5",
			OutputName: "chi",
		},
		{
			ImportPath: "github.com/jackc/pgx/v5",
			OutputName: "pgx",
		},
		{
			ImportPath: "github.com/nats-io/nats.go",
			OutputName: "nats-go",
		},
		{
			ImportPath: "github.com/gorilla/sessions",
			OutputName: "gorilla-sessions",
		},
	}

	contextDir := "context/go-packages"
	successCount := 0
	failCount := 0

	for _, pkg := range packages {
		outputFile := filepath.Join(contextDir, pkg.OutputName+".txt")

		slog.Info("Generating documentation", "package", pkg.ImportPath, "output", outputFile)

		// Run go doc -all for the package
		cmd := exec.Command("go", "doc", "-all", pkg.ImportPath)
		output, err := cmd.CombinedOutput()

		if err != nil {
			slog.Error("Failed to run go doc",
				"package", pkg.ImportPath,
				"error", err,
				"output", string(output))
			failCount++
			continue
		}

		// Write output to file
		if err := os.WriteFile(outputFile, output, 0644); err != nil {
			slog.Error("Failed to write file", "file", outputFile, "error", err)
			failCount++
			continue
		}

		successCount++
		slog.Info("Successfully generated documentation", "package", pkg.OutputName)
	}

	// Summary
	slog.Info("Documentation generation complete",
		"success", successCount,
		"failed", failCount,
		"total", len(packages))

	if failCount > 0 {
		fmt.Fprintf(os.Stderr, "\n%d package(s) failed to generate. See errors above.\n", failCount)
		os.Exit(1)
	}
}
