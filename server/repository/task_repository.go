package repository

import (
	"github.com/leoyakubov/go-crud-rest-api/server/model"
)

type ITaskRepository interface {
	Add(task *model.Task) error
	FindOneById(id int) (*model.Task, error)
	FindAll() ([]model.Task, error)
	UpdateById(id int, upd *model.Task) (*model.Task, error)
	DeleteById(id int) error
}

type TaskRepository struct {
	repo *Repository
}

func NewTaskRepository(r *Repository) ITaskRepository {
	return &TaskRepository{
		repo: r,
	}
}

func (tr *TaskRepository) Add(task *model.Task) error {
	var err = tr.repo.Add(&task)
	return err
}

func (tr *TaskRepository) FindOneById(id int) (*model.Task, error) {
	var (
		task model.Task
		err  error
	)
	err = tr.repo.FindOneById(&task, id)
	return &task, err
}

func (tr *TaskRepository) FindAll() ([]model.Task, error) {
	var (
		tasks []model.Task
		err   error
	)
	err = tr.repo.FindAll(&tasks)
	return tasks, err
}

func (tr *TaskRepository) UpdateById(id int, upd *model.Task) (*model.Task, error) {
	var (
		result model.Task
		err    error
	)
	err = tr.repo.Update(&upd, &result, id)
	return &result, err
}

func (tr *TaskRepository) DeleteById(id int) error {
	var (
		task model.Task
		err  error
	)
	err = tr.repo.Delete(&task, id)
	return err
}
