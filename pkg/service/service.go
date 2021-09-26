package service

import (
	"github.com/nikiandr/golang-todo-app"
	"github.com/nikiandr/golang-todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type List interface {
	CreateList(userId int, list todo.List) (int, error)
	GetAllLists(userId int) ([]todo.List, error)
	GetListById(userId, listId int) (todo.List, error)
	DeleteList(userId, listId int) error
	UpdateList(update todo.List, userId, listId int) error
}

type Item interface {
	CreateItem(userId, listId int, curItem todo.Item) (int, error)
	GetAllItems(userId, listId int) ([]todo.Item, error)
}

type Service struct {
	Authorization
	List
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		List:          NewTodoListService(repos),
		Item:          NewItemService(repos),
	}
}
