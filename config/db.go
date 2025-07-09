package config

import (
	"log"
	"os"

	"github.com/Koro-Erp/shared/models"
	"github.com/joho/godotenv"
)


func LoadDbConfig() models.DbConfig{
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found or couldn't load it — assuming Docker environment.")
    }

    return models.DbConfig{
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        DBSSLMode:  os.Getenv("DB_SSLMODE"),
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
	}
}