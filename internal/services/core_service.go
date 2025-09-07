package services

import "simple-todo-app/internal/repositories"

type Services struct {
	Task TaskService
	User UserService
}

func NewService(repo *repositories.Repositories) *Services {
	return &Services{
		Task: NewTaskService(repo.Task),
		User: NewUserService(repo.User),
	}
}
