package server

import (
	"net"

	"github.com/DevAthhh/todo-grpc/internal/entity"
	"github.com/DevAthhh/todo-grpc/internal/handler"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
)

type TodoService interface {
	Create(task *entity.Task) (*entity.Task, error)
	Delete(id int64) (*entity.Task, error)
	GetAll() []*entity.Task
	Update(task *entity.Task) (*entity.Task, error)
}

type Server struct {
	server *grpc.Server
}

func (s *Server) Run(port string) error {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	if err := s.server.Serve(l); err != nil {
		return err
	}

	return nil
}

func (s *Server) Shutdown() {
	s.server.GracefulStop()
}

func New(service TodoService) *Server {
	srv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))

	handler.Register(srv, service)

	return &Server{
		server: srv,
	}
}
