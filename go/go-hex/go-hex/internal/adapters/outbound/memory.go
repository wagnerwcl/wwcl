package outbound

import (
	"errors"
	"go-hex/internal/domain"
	"go-hex/internal/ports/outbound"
	"sync"
)

type InMemoryTaskRepository struct {
	tasks map[string]domain.Task
	mu    sync.RWMutex
}

func NewInMemoryTaskRepository() outbound.TaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]domain.Task),
	}
}

func (r *InMemoryTaskRepository) Save(task domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; exists {
		return errors.New("task with this ID already exists")
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}

func (r *InMemoryTaskRepository) FindAll() ([]domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]domain.Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *InMemoryTaskRepository) FindByID(id string) (domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return domain.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (r *InMemoryTaskRepository) Update(task domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return errors.New("task not found")
	}
	r.tasks[task.ID] = task
	return nil
}
