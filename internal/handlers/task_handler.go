package handlers

import (
	"encoding/json"
	"net/http"
	"simple-todo-app/internal/helpers"
	"simple-todo-app/internal/models"
	"strconv"
)

var (
	Task   = []models.Task{}
	nextId = 1
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	helpers.ApiResponse(w, 200, "Task loaded", Task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var data models.CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if len(data.Title) == 0 {
		helpers.ApiResponse(w, http.StatusUnprocessableEntity, "Title is required", nil)
		return
	}

	Task = append(Task, models.Task{
		Id:          nextId,
		Title:       data.Title,
		Description: data.Description,
		IsFinished:  false,
	})

	nextId++

	helpers.ApiResponse(w, 200, "New task created", data)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
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

	if len(data.Title) == 0 {
		helpers.ApiResponse(w, http.StatusUnprocessableEntity, "Title is required", nil)
		return
	}

	for idx, task := range Task {
		if task.Id == id {
			Task[idx].Title = data.Title
			Task[idx].Description = data.Description
			Task[idx].IsFinished = data.IsFinished

			helpers.ApiResponse(w, http.StatusOK, "Task updated", Task[idx])
			return
		}
	}

	helpers.ApiResponse(w, http.StatusNotFound, "Task not found", nil)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	for idx, task := range Task {
		if task.Id == id {
			Task = append(Task[:idx], Task[idx+1:]...)
			helpers.ApiResponse(w, http.StatusOK, "Task deleted", nil)
			return
		}
	}

	helpers.ApiResponse(w, http.StatusNotFound, "Task not found", nil)
}
