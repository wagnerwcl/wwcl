package service

import (
	"errors"
	"go-hex/internal/domain"
	"go-hex/internal/ports/inbound"
	"go-hex/internal/ports/outbound"
)

type TaskServiceImpl struct {
	repo outbound.TaskRepository
}

func NewTaskService(repo outbound.TaskRepository) inbound.TaskService {
	return &TaskServiceImpl{repo: repo}
}

func (s *TaskServiceImpl) AddTask(task domain.Task) error {
	if task.Description == "" {
		return errors.New("task description cannot be empty")
	}
	return s.repo.Save(task)
}

func (s *TaskServiceImpl) RemoveTask(id string) error {
	return s.repo.Delete(id)
}

func (s *TaskServiceImpl) ListTasks() ([]domain.Task, error) {
	return s.repo.FindAll()
}

func (s *TaskServiceImpl) MarkTaskAsCompleted(id string) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	task.Completed = true
	return s.repo.Update(task)
}
