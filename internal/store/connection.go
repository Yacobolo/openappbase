package store

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

const (
	// Default SQLite database path
	defaultDBPath = "./data/app.db"

	// Connection pool settings optimized for SQLite
	maxOpenConns    = 1 // SQLite works best with a single writer
	maxIdleConns    = 1
	connMaxLifetime = 1 * time.Hour
	connMaxIdleTime = 10 * time.Minute

	// Operation timeout
	defaultTimeout = 5 * time.Second
)

// Database wraps a SQL database connection with lifecycle management
type Database struct {
	*sql.DB
	path string
}

// InitDB initializes the SQLite database connection for the application's internal database.
// This is NOT for user's PostgreSQL databases - those are stored in the connections table.
//
// The database path can be overridden via SQLITE_DB_PATH environment variable.
// Default: ./data/app.db
//
// Returns error instead of calling os.Exit() to allow proper error handling and testing.
func InitDB(ctx context.Context) (*Database, error) {
	dbPath := getDBPath()

	// Ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	// Build connection string with SQLite pragmas as URL parameters (best practice)
	// Reference: https://github.com/mattn/go-sqlite3#connection-string
	connStr := buildConnectionString(dbPath)

	// Open SQLite database connection
	db, err := sql.Open("sqlite", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite database: %w", err)
	}

	// Configure connection pool settings
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	// Ping with context to verify the connection is alive
	pingCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		db.Close() // Clean up on failure
		return nil, fmt.Errorf("failed to ping SQLite database: %w", err)
	}

	slog.Info("successfully connected to SQLite database", "path", dbPath)

	return &Database{
		DB:   db,
		path: dbPath,
	}, nil
}

// Close closes the database connection gracefully
func (d *Database) Close() error {
	if d.DB == nil {
		return nil
	}

	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing database: %w", err)
	}

	slog.Info("database connection closed", "path", d.path)
	return nil
}

// Path returns the database file path
func (d *Database) Path() string {
	return d.path
}

// getDBPath returns the database path from environment or uses the default
func getDBPath() string {
	if path := os.Getenv("SQLITE_DB_PATH"); path != "" {
		return path
	}
	return defaultDBPath
}

// buildConnectionString constructs a SQLite connection string with pragmas as URL parameters
// This is the recommended approach over executing PRAGMA statements separately
func buildConnectionString(dbPath string) string {
	params := url.Values{}
	params.Add("_foreign_keys", "on")
	params.Add("_journal_mode", "WAL")
	params.Add("_synchronous", "NORMAL")
	params.Add("_busy_timeout", "5000")
	params.Add("_cache_size", "-64000") // 64MB cache

	return fmt.Sprintf("file:%s?%s", dbPath, params.Encode())
}
