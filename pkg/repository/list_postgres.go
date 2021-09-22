package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nikiandr/golang-todo-app"
	"github.com/pkg/errors"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{db: db}
}

func (r *ListPostgres) Create(userId int, list todo.List) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", listsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			return 0, errors.Errorf("Two errors with DB: №1 (%s) and №2 (%s)", err.Error(), errRoll.Error())
		}
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			return 0, errors.Errorf("Two errors with DB: №1 (%s) and №2 (%s)", err.Error(), errRoll.Error())
		}
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ListPostgres) GetAll(userId int) ([]todo.List, error) {
	var lists []todo.List

	query := fmt.Sprintf("SELECT t1.id, t1.title, t1.description FROM %s t1 INNER JOIN %s t2 ON t1.id = t2.list_id WHERE t2.user_id = $1",
		listsTable, usersListTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *ListPostgres) GetById(userId, listId int) (todo.List, error) {
	var list todo.List

	query := fmt.Sprintf("SELECT t1.id, t1.title, t1.description FROM %s t1 INNER JOIN %s t2 ON t1.id = t2.list_id WHERE t2.user_id = $1 AND t1.id = $2",
		listsTable, usersListTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}
