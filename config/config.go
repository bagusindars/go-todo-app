package config

type DatabaseConfig struct {
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
}

func LoadConfig() *DatabaseConfig {
	db := &DatabaseConfig{
		DB_PORT:     "5432",
		DB_HOST:     "localhost",
		DB_NAME:     "simple-todo-app",
		DB_USERNAME: "postgres",
		DB_PASSWORD: "",
	}

	return db
}
