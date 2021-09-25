package service

import (
	todo "github.com/nikiandr/golang-todo-app"
	"github.com/nikiandr/golang-todo-app/pkg/repository"
)

type ItemService struct {
	repo repository.Item
}

func (s *ItemService) CreateItem(userId, listId int, curItem todo.Item) (int, error) {
	return s.repo.CreateItem(userId, listId, curItem)
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}
