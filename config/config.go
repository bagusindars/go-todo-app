package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
}

func LoadConfig() *DatabaseConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error load ENV : " + err.Error())
	}

	db := &DatabaseConfig{
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}

	return db
}
