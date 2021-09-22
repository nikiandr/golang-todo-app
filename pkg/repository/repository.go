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
	Create(userId int, list todo.List) (int, error)
	GetAll(userId int) ([]todo.List, error)
	GetById(userId, listId int) (todo.List, error)
}

type Item interface {
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
	}
}
