package repository

import (
	"errors"
	"praktik-todo/internal/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *entity.Task) error
	FindAll() ([]entity.Task, error)
	FindByID(id uint) (*entity.Task, error)
	Update(task *entity.Task) error
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *entity.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) FindAll() ([]entity.Task, error) {
	var tasks []entity.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) FindByID(id uint) (*entity.Task, error) {
	var task entity.Task
	if err := r.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) Update(task *entity.Task) error {
	return r.db.Model(&entity.Task{}).Where("id = ?", task.ID).Updates(task).Error
}

func (r *taskRepository) Delete(id uint) error {
	if err := r.db.Delete(&entity.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
