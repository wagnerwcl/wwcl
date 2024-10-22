package main

import (
	"log"
	"net/http"
	"os"

	"go-hex/internal/adapters/inbound"
	"go-hex/internal/adapters/outbound"
	port "go-hex/internal/ports/outbound"
	"go-hex/internal/service"
)

const (
	AppTypeFile   = "file"
	AppTypeMemory = "memory"
)

var repo port.TaskRepository

func init() {
	appType := os.Getenv("APP_TYPE")

	switch appType {
	case AppTypeFile:
		log.Println("Using file repository")
		repoFile, err := outbound.NewFileTaskRepository("tasks.json")
		if err != nil {
			log.Fatalf("Failed to initialize repository: %s", err)
		}
		repo = repoFile
	case AppTypeMemory:
		log.Println("Using in-memory repository")
		repo = outbound.NewInMemoryTaskRepository()
	default:
		log.Fatalf("Invalid APP_TYPE: %s", appType)
	}
}

func main() {
	taskService := service.NewTaskService(repo)
	taskHandler := inbound.NewTaskHandler(taskService)

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			taskHandler.AddTask(w, r)
		case http.MethodGet:
			taskHandler.ListTasks(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/completed", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			taskHandler.MarkTaskAsCompleted(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			taskHandler.RemoveTask(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
