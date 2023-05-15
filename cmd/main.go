package main

import (
	ToDoRESTApi "github.com/MelomiLight/go-todo-app"
	"github.com/MelomiLight/go-todo-app/pkg/handler"
	"github.com/MelomiLight/go-todo-app/pkg/repository"
	service "github.com/MelomiLight/go-todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(ToDoRESTApi.Server)
	if err := srv.Run("8888", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
