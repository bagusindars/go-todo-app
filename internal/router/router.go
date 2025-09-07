package router

import (
	"log"
	"net/http"
	"simple-todo-app/internal/handlers"
)

func SetupRoute(handler *handlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/tasks", handler.Task.GetTask)
	mux.HandleFunc("POST /api/tasks", handler.Task.CreateTask)
	mux.HandleFunc("PUT /api/tasks/{id}", handler.Task.UpdateTask)
	mux.HandleFunc("DELETE /api/tasks/{id}", handler.Task.DeleteTask)

	mux.HandleFunc("POST /api/auth/register", handler.User.Register)
	mux.HandleFunc("POST /api/auth/login", handler.User.Login)

	return mux
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
