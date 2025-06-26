package main

import (
	"github.com/DevAthhh/todo-grpc/internal/repository"
	"github.com/DevAthhh/todo-grpc/internal/server"
	"github.com/DevAthhh/todo-grpc/internal/service"
	"github.com/DevAthhh/todo-grpc/pkg/config"
	"github.com/DevAthhh/todo-grpc/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	cfg := config.MustLoad()

	log, err := logger.Load(cfg.Env)
	if err != nil {
		panic(err)
	}

	repo := repository.New()
	serv := service.New(repo)
	srv := server.New(serv)

	log.Info("server running...")
	if err := srv.Run(cfg.Server.Port); err != nil {
		log.Error(err.Error())
		return
	}

}
