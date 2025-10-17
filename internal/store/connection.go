// store/connection.go
package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// db is a package-level variable to hold the database connection pool.
var db *sql.DB

// InitDB initializes the database connection using the provided connection string.
func InitDB(connStr string) {
	var err error
	// Open a connection to the database.
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Ping the database to verify the connection is alive.
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v\n", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
}

// CloseDB closes the database connection. It's good practice to call this
// when the application is shutting down.
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
