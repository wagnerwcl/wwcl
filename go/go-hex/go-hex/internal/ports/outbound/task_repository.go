package outbound

import "go-hex/internal/domain"

type TaskRepository interface {
	Save(task domain.Task) error
	Delete(id string) error
	FindAll() ([]domain.Task, error)
	FindByID(id string) (domain.Task, error)
	Update(task domain.Task) error
}
