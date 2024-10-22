package inbound

import (
	"encoding/json"
	"go-hex/internal/domain"
	"go-hex/internal/ports/inbound"
	"net/http"

	"github.com/google/uuid"
)

type TaskHandler struct {
	service inbound.TaskService
}

func NewTaskHandler(service inbound.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	task.ID = uuid.New().String()
	task.Completed = false

	if err := h.service.AddTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) RemoveTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Missing task ID", http.StatusBadRequest)
		return
	}

	if err := h.service.RemoveTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.ListTasks()

	if err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) MarkTaskAsCompleted(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing task ID", http.StatusBadRequest)
		return
	}

	if err := h.service.MarkTaskAsCompleted(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
