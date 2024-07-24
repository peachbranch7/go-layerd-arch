package infra

import (
	"go-layered-arch/domain/model"
	"go-layered-arch/domain/repository"

	"github.com/jinzhu/gorm"
)

type TaskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{Conn: conn}
}

func (tr *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Create(&task).Error; err != nil {
		return nil, err
	}

	resTask := &model.Task{
		ID: task.ID,
		Title: task.Title,
		Content: task.Content,
	}

	return resTask, nil
}

func (tr *TaskRepository) FindById(id int) (*model.Task, error) {
	task := &model.Task{
		ID: id,
	}

	if err := tr.Conn.Find(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (tr * TaskRepository) Update(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Model(&task).Updates(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) Delete(task *model.Task) error {
	if err := tr.Conn.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}

