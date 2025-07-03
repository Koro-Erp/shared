package db

import (
	"database/sql"
	"fmt"
	"log"
    _ "github.com/lib/pq"
)

func Connect(dbUsername string, dbPassword string, dbName string,dbSSLMode string) *sql.DB {
    connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
        dbUsername, dbPassword, dbName, dbSSLMode)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }
    return db
}
