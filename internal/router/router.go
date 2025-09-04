package router

import (
	"net/http"
	"simple-todo-app/internal/handlers"
)

func SetupRoute(handler *handlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/tasks", handler.Task.GetTask)
	mux.HandleFunc("POST /api/tasks", handler.Task.CreateTask)
	mux.HandleFunc("PUT /api/tasks/{id}", handler.Task.UpdateTask)
	mux.HandleFunc("DELETE /api/tasks/{id}", handler.Task.DeleteTask)

	return mux
}
