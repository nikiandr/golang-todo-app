package repository

import (
	"fmt"
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
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, username, password_hash) VALUES ($1, $2, $3, $4) RETURNING id",
		usersTable)

	row := r.db.QueryRow(query, user.Name, user.Surname, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthSql) GetUser(username string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id, password_hash from %s WHERE username=$1", usersTable)
	err := r.db.Get(&user, query, username)
	return user, err
}
