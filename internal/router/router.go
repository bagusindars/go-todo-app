package router

import (
	"net/http"
	"simple-todo-app/internal/handlers"
	"simple-todo-app/middleware"
)

func CustomMux() {

}

func SetupRoute(handler *handlers.Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/auth/register", handler.User.Register)
	mux.HandleFunc("POST /api/auth/login", handler.User.Login)

	middleware.HandlerWithMiddleware(mux, "GET /api/tasks", handler.Task.GetTask, middleware.AuthJWTMiddleware)
	middleware.HandlerWithMiddleware(mux, "POST /api/tasks", handler.Task.CreateTask, middleware.AuthJWTMiddleware)
	middleware.HandlerWithMiddleware(mux, "PUT /api/tasks/{id}", handler.Task.UpdateTask, middleware.AuthJWTMiddleware)
	middleware.HandlerWithMiddleware(mux, "DELETE /api/tasks/{id}", handler.Task.DeleteTask, middleware.AuthJWTMiddleware)

	middleware.HandlerWithMiddleware(mux, "GET /api/profile", handler.User.Info, middleware.AuthJWTMiddleware)

	return middleware.LoggingMiddleware(mux)
}
