package services

import (
	"errors"
	"simple-todo-app/internal/helpers"
	"simple-todo-app/internal/models"
	"simple-todo-app/internal/repositories"
)

type TaskService interface {
	GetByUser(userId uint) ([]models.Task, error)
	CreateTask(userInfo *helpers.Claims, data models.Task) error
	UpdateTask(id int, userInfo *helpers.Claims, data models.Task) error
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

func (s *taskService) CreateTask(userInfo *helpers.Claims, data models.Task) error {
	if len(data.Title) == 0 {
		return errors.New("title cannot be empty")
	}

	err := s.repo.Create(data)

	if err != nil {
		return err
	}

	go helpers.SendEmail(userInfo.Email, "New task added", "You created new task : "+data.Title+". Focus on it!")

	return nil
}

func (s *taskService) UpdateTask(id int, userInfo *helpers.Claims, data models.Task) error {
	if len(data.Title) == 0 {
		return errors.New("title cannot be empty")
	}

	err := s.repo.Update(id, data)

	if err != nil {
		return err
	}

	if data.IsFinished {
		go helpers.SendEmail(
			userInfo.Email,
			data.Title+" is finished!",
			"Nice work! You’ve just finished the task: "+data.Title+". Keep up the great momentum 🚀",
		)
	}

	return nil
}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
