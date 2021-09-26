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

func (s *ItemService) GetItemById(userId, listId, itemId int) (todo.Item, error) {
	return s.repo.GetItemById(userId, listId, itemId)
}

func (s *ItemService) UpdateItem(update todo.Item, userId, listId, itemId int) error {
	return s.repo.UpdateItem(update, userId, listId, itemId)
}

func (s *ItemService) DeleteItem(userId, listId, itemId int) error {
	return s.repo.DeleteItem(userId, listId, itemId)
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}
