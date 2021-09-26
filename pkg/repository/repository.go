package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nikiandr/golang-todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string) (todo.User, error)
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
	GetItemById(userId, listId, itemId int) (todo.Item, error)
	UpdateItem(update todo.Item, userId, listId, itemId int) error
	DeleteItem(userId, listId, itemId int) error
}

type Repository struct {
	Authorization
	List
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		List:          NewListPostgres(db),
		Item:          NewItemPostgres(db),
	}
}
