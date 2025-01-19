package usecase

import (
	"praktik-todo/internal/entity"
	"praktik-todo/internal/repository"
)

type TaskUsecase interface {
	CreateTask(task *entity.Task) error
	GetAllTasks() ([]entity.Task, error)
	GetTaskByID(id uint) (*entity.Task, error)
	UpdateTask(task *entity.Task) error
	DeleteTask(id uint) error
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (u *taskUsecase) CreateTask(task *entity.Task) error {
	return u.taskRepo.Create(task)
}

func (u *taskUsecase) GetAllTasks() ([]entity.Task, error) {
	return u.taskRepo.FindAll()
}

func (u *taskUsecase) GetTaskByID(id uint) (*entity.Task, error) {
	return u.taskRepo.FindByID(id)
}

func (u *taskUsecase) UpdateTask(task *entity.Task) error {
	return u.taskRepo.Update(task)
}

func (u *taskUsecase) DeleteTask(id uint) error {
	return u.taskRepo.Delete(id)
}
