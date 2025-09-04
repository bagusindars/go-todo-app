package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simple-todo-app/db"
	"simple-todo-app/internal/config"
	"simple-todo-app/internal/handlers"
	"simple-todo-app/internal/repositories"
	"simple-todo-app/internal/router"
	"simple-todo-app/internal/services"
	"syscall"
	"time"
)

type Application struct {
	config  *config.Config
	db      *db.DB
	handler *handlers.Handlers
}

func main() {
	// load config
	cfg := config.Load()

	app := NewApplication(cfg)

	defer app.Close()

	router := router.SetupRoute(app.handler)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler: router,
	}

	go func() {
		fmt.Println("Starting server port :", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}

func NewApplication(cfg *config.Config) *Application {
	// database & migration
	connection := db.InitDatabase(cfg.Database)
	migrationCon := db.InitDatabase(cfg.Database)
	migration := db.NewMigrator(migrationCon.DB)
	if err := migration.Up(); err != nil {
		fmt.Println(err.Error())
	}

	// repository
	repos := repositories.NewRepositories(connection.DB)

	// service
	svc := services.NewService(repos)

	// handler
	hs := handlers.NewHandlers(svc)

	return &Application{
		config:  cfg,
		db:      connection,
		handler: hs,
	}
}

func (a *Application) Close() {
	if a.db != nil {
		a.db.Close()
	}
}
