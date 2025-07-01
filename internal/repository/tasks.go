package repository

import "github.com/Rafiur/api-crud-assessment/internal/model"

type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	GetByID(id int) (*model.Task, error)
	List() ([]model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(id int) error
}
