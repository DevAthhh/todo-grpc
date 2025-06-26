package handler

import (
	"context"
	"errors"

	todo "github.com/DevAthhh/protos-todo/gen/go"
	"github.com/DevAthhh/todo-grpc/internal/entity"
	"google.golang.org/grpc"
)

type TodoService interface {
	Create(task *entity.Task) (*entity.Task, error)
	Delete(id int64) (*entity.Task, error)
	GetAll() []*entity.Task
	Update(task *entity.Task) (*entity.Task, error)
}

type ServerAPI struct {
	todo.UnimplementedTodoServer
	todo TodoService
}

func Register(server *grpc.Server, todoService TodoService) {
	todo.RegisterTodoServer(server, &ServerAPI{todo: todoService})
}

func (s *ServerAPI) Create(ctx context.Context, req *todo.CreateRequest) (*todo.CreateResponse, error) {
	if req.GetTitle() == "" {
		return nil, errors.New("title is empty")
	}

	var task *entity.Task
	var err error

	task = &entity.Task{
		Title: req.GetTitle(),
		Desc:  req.GetDescription(),
	}

	task, err = s.todo.Create(task)
	if err != nil {
		return nil, errors.New("err with creating task")
	}

	return &todo.CreateResponse{
		Title:       task.Title,
		Description: task.Desc,
		Id:          task.ID,
		Status:      task.Status,
	}, nil
}

func (s *ServerAPI) Delete(ctx context.Context, req *todo.DeleteRequest) (*todo.DeleteResponse, error) {
	if req.GetId() == 0 {
		return nil, errors.New("invalid id")
	}

	task, err := s.todo.Delete(req.GetId())
	if err != nil {
		return nil, errors.New("unknown task")
	}

	return &todo.DeleteResponse{
		Title:       task.Title,
		Description: task.Desc,
		Id:          task.ID,
		Status:      task.Status,
	}, nil
}

func (s *ServerAPI) GetAll(ctx context.Context, req *todo.GetAllRequest) (*todo.GetAllResponse, error) {
	task := s.todo.GetAll()

	tasks := make([]*todo.Tasks, 0)
	for _, t := range task {
		tasks = append(tasks, &todo.Tasks{
			Title:       t.Title,
			Description: t.Desc,
			Id:          t.ID,
			Status:      t.Status,
		})
	}

	return &todo.GetAllResponse{
		Task: tasks,
	}, nil
}

func (s *ServerAPI) UpdateOne(ctx context.Context, req *todo.UpdateRequest) (*todo.UpdateResponse, error) {
	if req.GetId() == 0 {
		return nil, errors.New("invalid id")
	}

	task := entity.Task{
		Status: req.GetStatus(),
		ID:     req.GetId(),
	}

	updTask, err := s.todo.Update(&task)
	if err != nil {
		return nil, errors.New("unknown task")
	}

	return &todo.UpdateResponse{
		Title:       updTask.Title,
		Description: updTask.Desc,
		Status:      updTask.Status,
		Id:          updTask.ID,
	}, nil
}
