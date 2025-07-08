package db

import (
	"database/sql"
	"fmt"
	"log"
    _ "github.com/lib/pq"
)

func Connect(dbHost string, dbPort string, dbUsername string, dbPassword string, dbName string, dbSSLMode string) *sql.DB {
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        dbHost, dbPort, dbUsername, dbPassword, dbName, dbSSLMode,
    )

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }

    return db
}
