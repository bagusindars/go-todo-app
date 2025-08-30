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

	connectionString := "postgresql://" + dbConfig.DB_USERNAME + ":" + dbConfig.DB_PASSWORD + "@localhost:" + dbConfig.DB_PORT + "/" + dbConfig.DB_NAME + "?sslmode=disable"

	var err error

	connection, err = sql.Open("postgres", connectionString)

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
