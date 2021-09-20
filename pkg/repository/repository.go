package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nikiandr/golang-todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type List interface {
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
		Authorization: NewAuthSql(db),
	}
}
