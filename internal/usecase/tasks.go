package usecase

import (
	"github.com/Rafiur/api-crud-assessment/internal/model"
	"github.com/Rafiur/api-crud-assessment/internal/repository"
)

type TaskService struct {
	TaskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) *TaskService {
	return &TaskService{TaskRepo: taskRepo}
}

func (service *TaskService) Create(task *model.Task) (*model.Task, error) {
	return service.TaskRepo.Create(task)
}

func (service *TaskService) GetById(id int) (*model.Task, error) {
	return service.TaskRepo.GetByID(id)
}

func (service *TaskService) List() ([]model.Task, error) {
	return service.TaskRepo.List()
}

func (service *TaskService) Update(task *model.Task) (*model.Task, error) {
	return service.TaskRepo.Update(task)
}

func (service *TaskService) Delete(id int) error {
	return service.TaskRepo.Delete(id)
}
