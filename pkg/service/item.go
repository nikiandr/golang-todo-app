package service

import (
	"github.com/nikiandr/golang-todo-app"
	"github.com/nikiandr/golang-todo-app/pkg/repository"
)

type ItemService struct {
	repo repository.Item
}

func (s *ItemService) CreateItem(userId, listId int, curItem todo.Item) (int, error) {
	return s.repo.CreateItem(userId, listId, curItem)
}

func (s *ItemService) GetAllItems(userId, listId int) ([]todo.Item, error) {
	return s.repo.GetAllItems(userId, listId)
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}
