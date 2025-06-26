package repository

import (
	"errors"

	"github.com/DevAthhh/todo-grpc/internal/entity"
)

type TodoRepository struct {
	tasks []*entity.Task
}

func (r *TodoRepository) Create(task *entity.Task) (*entity.Task, error) {
	if len(r.tasks) == 0 {
		task.ID = 1
	} else {
		task.ID = r.tasks[len(r.tasks)-1].ID + 1
	}

	task.Status = "in process"

	r.tasks = append(r.tasks, task)

	return task, nil
}

func (r *TodoRepository) Delete(id int64) (*entity.Task, error) {
	var task *entity.Task

	for idx := range r.tasks {
		if r.tasks[idx].ID == id {
			task = r.tasks[idx]
			tmp := r.tasks[:idx]
			r.tasks = r.tasks[idx+1:]
			r.tasks = append(r.tasks, tmp...)
			break
		}
	}

	if task == nil {
		return nil, errors.New("unknown task")
	}

	return task, nil
}

func (r *TodoRepository) GetAll() []*entity.Task {
	return r.tasks
}

func (r *TodoRepository) Update(task *entity.Task) (*entity.Task, error) {
	var taskRes *entity.Task

	for idx := range r.tasks {
		if r.tasks[idx].ID == task.ID {
			r.tasks[idx].Status = task.Status
			taskRes = r.tasks[idx]
			break
		}
	}

	if taskRes == nil {
		return nil, errors.New("unknown task")
	}

	return taskRes, nil
}

func New() *TodoRepository {
	return &TodoRepository{
		tasks: make([]*entity.Task, 0),
	}
}
