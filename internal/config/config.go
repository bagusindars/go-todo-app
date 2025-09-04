package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Port     string
	Host     string
	DBName   string
	User     string
	Password string
	Address  string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load ENV : " + err.Error())
	}

	connectionString := "postgresql://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"

	return &Config{
		Server: ServerConfig{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: DatabaseConfig{
			Port:     os.Getenv("DB_PORT"),
			Host:     os.Getenv("DB_HOST"),
			DBName:   os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Address:  connectionString,
		},
	}
}
