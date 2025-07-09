package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Koro-Erp/shared/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func Connect(dbConfig models.DbConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBSSLMode,
	)

	// log.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Verify the connection is actually working
	// err = db.Ping()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to ping database: %w", err)
	// }

	// Configure connection pool settings (recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}