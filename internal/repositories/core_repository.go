package repositories

import "database/sql"

type Repositories struct {
	Task TaskRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Task: NewTaskRepository(db),
	}
}
