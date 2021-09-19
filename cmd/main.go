package main

import (
	"github.com/nikiandr/golang-todo-app"
	"github.com/nikiandr/golang-todo-app/pkg/handler"
	"github.com/nikiandr/golang-todo-app/pkg/repository"
	"github.com/nikiandr/golang-todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
