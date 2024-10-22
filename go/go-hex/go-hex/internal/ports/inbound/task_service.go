package inbound

import "go-hex/internal/domain"

type TaskService interface {
	AddTask(task domain.Task) error
	RemoveTask(id string) error
	ListTasks() ([]domain.Task, error)
	MarkTaskAsCompleted(id string) error
}
