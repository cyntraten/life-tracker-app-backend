package tasks

import "gorm.io/gorm"

type TasksRepository interface {
	GetAllTasks() ([]Task, error)
	CreateTask(task Task) error
	DeleteTask(id string) error
	UpdateTask(task Task) error
	GetTaskById(id string) (Task, error)
}

type tasksRepo struct {
	db *gorm.DB
}

func NewTasksRepository(db *gorm.DB) TasksRepository {
	return &tasksRepo{db: db}
}

func (r *tasksRepo) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *tasksRepo) CreateTask(task Task) error {
	return r.db.Create(&task).Error
}

func (r *tasksRepo) DeleteTask(id string) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}

func (r *tasksRepo) UpdateTask(task Task) error {
	return r.db.Save(&task).Error
}

func (r *tasksRepo) GetTaskById(id string) (Task, error) {
	var task Task

	err := r.db.First(&Task{}, "id = ?", id).Error

	return task, err
}
