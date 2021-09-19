package main

import (
	"github.com/nikiandr/golang-todo-app"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
