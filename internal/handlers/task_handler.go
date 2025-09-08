package handlers

import (
	"encoding/json"
	"net/http"
	"simple-todo-app/internal/helpers"
	"simple-todo-app/internal/models"
	"simple-todo-app/internal/services"
	"strconv"
)

type TaskHandler struct {
	service services.TaskService
}

func NewTaskHandler(s services.TaskService) *TaskHandler {
	return &TaskHandler{
		service: s,
	}
}

func (s *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*helpers.Claims)
	tasks, err := s.service.GetByUser(userInfo.UserId)

	if err != nil {
		helpers.ApiResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.ApiResponse(w, http.StatusOK, "Task loaded", tasks)
}

func (s *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var data models.CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userInfo := r.Context().Value("userInfo").(*helpers.Claims)

	task := models.Task{
		Title:       data.Title,
		Description: data.Description,
		IsFinished:  false,
		UserId:      int(userInfo.UserId),
	}

	err = s.service.CreateTask(userInfo, task)

	if err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.ApiResponse(w, http.StatusOK, "Task created", task)
}

func (s *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	var data models.UpdateTaskRequest

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userInfo := r.Context().Value("userInfo").(*helpers.Claims)

	task := models.Task{
		Id:          id,
		Title:       data.Title,
		Description: data.Description,
		IsFinished:  data.IsFinished,
		UserId:      int(userInfo.UserId),
	}

	if err = s.service.UpdateTask(id, userInfo, task); err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.ApiResponse(w, http.StatusOK, "Task updated", task)
}

func (s *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	if err = s.service.DeleteTask(id); err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.ApiResponse(w, http.StatusOK, "Task deleted", nil)
}
