package main

import (
	"Todolist"
	"Todolist/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(Todolist.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
