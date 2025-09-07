package repositories

import "database/sql"

type Repositories struct {
	Task TaskRepository
	User UserRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Task: NewTaskRepository(db),
		User: NewUserRepository(db),
	}
}
