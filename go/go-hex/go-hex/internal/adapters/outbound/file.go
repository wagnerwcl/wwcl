package outbound

import (
	"encoding/json"
	"errors"
	"go-hex/internal/domain"
	"go-hex/internal/ports/outbound"
	"os"
	"sync"
)

type FileTaskRepository struct {
	filePath string
	tasks    map[string]domain.Task
	mu       sync.RWMutex
}

func NewFileTaskRepository(filePath string) (outbound.TaskRepository, error) {
	repo := &FileTaskRepository{
		filePath: filePath,
		tasks:    make(map[string]domain.Task),
	}

	if err := repo.load(); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *FileTaskRepository) load() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		r.tasks = make(map[string]domain.Task)
		return nil
	}

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		r.tasks = make(map[string]domain.Task)
		return nil
	}

	if err := json.Unmarshal(data, &r.tasks); err != nil {
		return err
	}
	return nil
}

func (r *FileTaskRepository) saveToFile() error {
	data, err := json.MarshalIndent(r.tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}

func (r *FileTaskRepository) Save(task domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; exists {
		return errors.New("task with this ID already exists")
	}
	r.tasks[task.ID] = task
	return r.saveToFile()
}

func (r *FileTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return r.saveToFile()
}

func (r *FileTaskRepository) FindAll() ([]domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]domain.Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *FileTaskRepository) FindByID(id string) (domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]

	if !exists {
		return domain.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (r *FileTaskRepository) Update(task domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return errors.New("task not found")
	}
	r.tasks[task.ID] = task
	return r.saveToFile()
}
