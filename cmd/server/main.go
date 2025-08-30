package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-todo-app/db"
	"simple-todo-app/internal/handlers"
)

func main() {
	port := "9000"

	// database
	db.Init()
	defer db.Close()

	// route
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/tasks", handlers.GetTask)
	mux.HandleFunc("POST /api/tasks", handlers.CreateTask)
	mux.HandleFunc("PUT /api/tasks/{id}", handlers.UpdateTask)
	mux.HandleFunc("DELETE /api/tasks/{id}", handlers.DeleteTask)

	fmt.Println("Server started on port :", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
