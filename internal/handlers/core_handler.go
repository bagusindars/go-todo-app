package handlers

import "simple-todo-app/internal/services"

type Handlers struct {
	Task *TaskHandler
	User *UserHandler
}

func NewHandlers(svc *services.Services) *Handlers {
	return &Handlers{
		Task: NewTaskHandler(svc.Task),
		User: NewUserHandler(svc.User),
	}
}
