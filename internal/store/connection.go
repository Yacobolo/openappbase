package store

import (
	"context" // pgx v5 requires a context for all operations
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// pool is a package-level variable to hold the database connection pool.
// We use *pgxpool.Pool instead of *sql.DB.
var pool *pgxpool.Pool

// InitDB initializes the database connection using the provided connection string.
// It now returns a *pgxpool.Pool.
func InitDB(connStr string) *pgxpool.Pool {
	var err error

	// pgxpool.New is used to create a new connection pool.
	// It's the equivalent of sql.Open for the native pgx driver.
	pool, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		// Using fmt.Fprintf to write to stderr and os.Exit to terminate.
		// This is often preferred over log.Fatalf for more control.
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	// Ping the database to verify the connection is alive.
	// The pool.Ping() method requires a context.
	if err := pool.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to ping database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to PostgreSQL using pgxpool!")
	return pool
}

// CloseDB closes all connections in the pool.
func CloseDB() {
	if pool != nil {
		pool.Close()
		fmt.Println("Database connection pool closed.")
	}
}
