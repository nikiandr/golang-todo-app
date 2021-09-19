package main

import (
	"github.com/nikiandr/golang-todo-app"
	"github.com/nikiandr/golang-todo-app/pkg/handler"
	"github.com/nikiandr/golang-todo-app/pkg/repository"
	"github.com/nikiandr/golang-todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error reading configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
