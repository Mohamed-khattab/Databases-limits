package configs

import (
	"github.com/joho/Godotenv"
	"log"
	"os"
)


func LoadEnv() {
	err := Godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type configs struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

func newConfigs() *configs {
	return &configs{
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),
	}
}