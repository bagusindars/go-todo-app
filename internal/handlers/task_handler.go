package handlers

import (
	"encoding/json"
	"net/http"
	"simple-todo-app/db"
	"simple-todo-app/internal/helpers"
	"simple-todo-app/internal/models"
	"strconv"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Connection().Query("SELECT id, title, description, is_finished from tasks")

	if err != nil {
		helpers.ApiResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	defer rows.Close()

	var tasks = []models.Task{}
	for rows.Next() {
		var each models.Task
		if err := rows.Scan(&each.Id, &each.Title, &each.Description, &each.IsFinished); err != nil {
			helpers.ApiResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		tasks = append(tasks, each)
	}

	if err = rows.Err(); err != nil {
		helpers.ApiResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.ApiResponse(w, 200, "Task loaded", tasks)
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

	_, err = db.Connection().Exec("INSERT INTO tasks (title, description) values ($1, $2)", data.Title, data.Description)

	if err != nil {
		helpers.ApiResponse(w, http.StatusInternalServerError, "Error Insert Task : "+err.Error(), nil)
		return
	}

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

	res, err := db.Connection().Exec("UPDATE tasks SET title = $1, description = $2, is_finished = $3 WHERE id = $4", data.Title, data.Description, data.IsFinished, id)

	if err != nil {
		helpers.ApiResponse(w, http.StatusInternalServerError, "Error update task : "+err.Error(), nil)
		return
	}

	rows, _ := res.RowsAffected()

	if rows == 0 {
		helpers.ApiResponse(w, http.StatusNotFound, "Task not found", nil)
		return
	}

	helpers.ApiResponse(w, http.StatusOK, "Task updated", nil)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	res, err := db.Connection().Exec("DELETE from tasks where id = $1", id)

	if err != nil {
		helpers.ApiResponse(w, http.StatusInternalServerError, "Error delete task : "+err.Error(), nil)
		return
	}

	// optional. just check if data with id is exists
	rows, _ := res.RowsAffected()

	if rows == 0 {
		helpers.ApiResponse(w, http.StatusNotFound, "Task not found", nil)
		return
	}

	helpers.ApiResponse(w, http.StatusOK, "Task deleted", nil)
}
