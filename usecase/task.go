package usecase

import (
	"go-layered-arch/domain/model"
	"go-layered-arch/domain/repository"
)

type TaskUsecase interface {
	Create(title, content string) (*model.Task, error)
	FindById(id int) (*model.Task, error)
	Update(id int, title, content string) (*model.Task, error)
	Delete(id int) error
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (tu *taskUsecase)  Create(title, content string) (*model.Task, error) {
	task, err := model.NewTask(title, content)

	if err != nil {
		return nil, err
	}

	createTask, err := tu.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return createTask, nil
}

func (tu *taskUsecase) FindById(id int) (*model.Task, error) {
	task, err := tu.taskRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tu *taskUsecase) Update(id int, title, content string) (*model.Task, error) {
	task, err := tu.taskRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	err = task.Set(title, content)
	if err != nil {
		return nil, err
	}

	updatedTask, err := tu.taskRepo.Update(task)

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (tu *taskUsecase) Delete(id int) error {
	task, err := tu.taskRepo.FindById(id)
	if err != nil {
		return err
	}

	err = tu.taskRepo.Delete(task)

	if err != nil {
		return err
	}

	return nil
}

