// configs/configs.go

package configs

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

// Config holds the database configuration
type Config struct {
    DbHost string
    DbPort string
    DbUser string
    DbPass string
    DbName string
}

// NewConfig loads the configuration from .env file and returns a new Config instance
func NewConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    return &Config{
        DbHost: os.Getenv("DB_HOST"),
        DbPort: os.Getenv("DB_PORT"),
        DbUser: os.Getenv("DB_USER"),
        DbPass: os.Getenv("DB_PASS"),
        DbName: os.Getenv("DB_NAME"),
    }
}
