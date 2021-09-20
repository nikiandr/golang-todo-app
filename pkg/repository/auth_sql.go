package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/nikiandr/golang-todo-app"
)

type AuthSql struct {
	db *sqlx.DB
}

func NewAuthSql(db *sqlx.DB) *AuthSql {
	return &AuthSql{db: db}
}

func (r *AuthSql) CreateUser(user todo.User) (int, error) {
	return 0, nil
}
