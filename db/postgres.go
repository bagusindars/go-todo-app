package db

import (
	"database/sql"
	"log"
	"simple-todo-app/internal/config"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func InitDatabase(cfg config.DatabaseConfig) *DB {
	connection, err := sql.Open("postgres", cfg.Address)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &DB{connection}
}
