package service

import "github.com/DevAthhh/todo-grpc/internal/entity"

type TodoRepo interface {
	Create(task *entity.Task) (*entity.Task, error)
	Delete(id int64) (*entity.Task, error)
	GetAll() []*entity.Task
	Update(task *entity.Task) (*entity.Task, error)
}

type todoService struct {
	repo TodoRepo
}

func (ts *todoService) Create(task *entity.Task) (*entity.Task, error) {
	return ts.repo.Create(task)
}

func (ts *todoService) Delete(id int64) (*entity.Task, error) {
	return ts.repo.Delete(id)
}

func (ts *todoService) GetAll() []*entity.Task {
	return ts.repo.GetAll()
}

func (ts *todoService) Update(task *entity.Task) (*entity.Task, error) {
	return ts.repo.Update(task)
}

func New(repo TodoRepo) *todoService {
	return &todoService{
		repo: repo,
	}
}
