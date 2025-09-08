package services

import (
	"errors"
	"simple-todo-app/internal/models"
	"simple-todo-app/internal/repositories"
)

type TaskService interface {
	GetByUser(userId uint) ([]models.Task, error)
	CreateTask(data models.Task) error
	UpdateTask(id int, data models.Task) error
	DeleteTask(id int) error
}

type taskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) GetByUser(userId uint) ([]models.Task, error) {
	return s.repo.GetAllByUser(userId)
}

func (s *taskService) CreateTask(data models.Task) error {
	if len(data.Title) == 0 {
		return errors.New("title cannot be empty")
	}

	return s.repo.Create(data)
}

func (s *taskService) UpdateTask(id int, data models.Task) error {
	if len(data.Title) == 0 {
		return errors.New("title cannot be empty")
	}

	return s.repo.Update(id, data)
}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
