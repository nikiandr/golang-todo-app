package service

import (
	"github.com/nikiandr/golang-todo-app"
	"github.com/nikiandr/golang-todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.List
}

func (s *TodoListService) CreateList(userId int, list todo.List) (int, error) {
	return s.repo.CreateList(userId, list)
}

func (s *TodoListService) GetAllLists(userId int) ([]todo.List, error) {
	return s.repo.GetAllLists(userId)
}

func (s *TodoListService) GetListById(userId, listId int) (todo.List, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *TodoListService) DeleteList(userId, listId int) error {
	return s.repo.DeleteList(userId, listId)
}

func (s *TodoListService) UpdateList(update todo.List, userId, listId int) error {
	return s.repo.UpdateList(update, userId, listId)
}

func NewTodoListService(repo repository.List) *TodoListService {
	return &TodoListService{repo: repo}
}
