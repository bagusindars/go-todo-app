package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
	conn *sql.DB
}

func NewMigrator(db *sql.DB) *Migrator {
	return &Migrator{
		conn: db,
	}
}

func (m *Migrator) CreateMigration() (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(m.conn, &postgres.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to create postgres driver: %w", err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create migration instance: %w", err)
	}

	return migration, nil
}

func (m *Migrator) Up() error {
	migration, err := m.CreateMigration()

	if err != nil {
		return err
	}

	defer migration.Close()

	err = migration.Up()

	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run up migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}

func (m *Migrator) Down() error {
	migration, err := m.CreateMigration()

	if err != nil {
		return err
	}

	defer migration.Close()

	err = migration.Down()

	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run down migrations: %w", err)
	}

	log.Println("Migrations rolled back successfully")
	return nil
}
