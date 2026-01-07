package main

import (
	"log"

	"Todolist"
	"Todolist/pkg/handler"
	"Todolist/pkg/repository"
	"Todolist/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)

	handlers := handler.NewHandler(services)
	srv := new(Todolist.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
