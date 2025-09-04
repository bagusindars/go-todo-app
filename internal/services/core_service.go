package services

import "simple-todo-app/internal/repositories"

type Services struct {
	Task taskService
}

func NewService(repo *repositories.Repositories) *Services {
	return &Services{
		Task: NewTaskService(repo.Task),
	}
}
