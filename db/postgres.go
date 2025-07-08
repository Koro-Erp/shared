package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	
	_ "github.com/lib/pq" // PostgreSQL driver
)

func Connect(dbHost, dbPort, dbUsername, dbPassword, dbName, dbSSLMode string) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUsername, dbPassword, dbName, dbSSLMode,
	)

	log.Println("Connecting to database...") // Less sensitive than logging full connection string
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Verify the connection is actually working
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Configure connection pool settings (recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}