package service

import (
	"github.com/nikiandr/golang-todo-app"
	"github.com/nikiandr/golang-todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.List
}

func (s *TodoListService) Create(userId int, list todo.List) (int, error) {
	return s.repo.Create(userId, list)
}

func NewTodoListService(repo repository.List) *TodoListService {
	return &TodoListService{repo: repo}
}
