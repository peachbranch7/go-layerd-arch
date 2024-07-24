package repository

import "go-layered-arch/domain/model"

type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	FindById(id int) (*model.Task, error)
	Update(*model.Task) (*model.Task, error)
	Delete(*model.Task) error
}