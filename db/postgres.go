package db

import (
	"database/sql"
	"log"
	"simple-todo-app/config"

	_ "github.com/lib/pq"
)

var connection *sql.DB

func Init() {
	dbConfig := config.LoadConfig()

	var err error

	connection, err = sql.Open("postgres", dbConfig.DB_DNS)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func Connection() *sql.DB {
	return connection
}

func Close() {
	if connection != nil {
		connection.Close()
	}
}
