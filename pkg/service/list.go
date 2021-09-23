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

func (s *TodoListService) GetAll(userId int) ([]todo.List, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (todo.List, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(update todo.List, userId, listId int) error {
	return s.repo.Update(update, userId, listId)
}

func NewTodoListService(repo repository.List) *TodoListService {
	return &TodoListService{repo: repo}
}
