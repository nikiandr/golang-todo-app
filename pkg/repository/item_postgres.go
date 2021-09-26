package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nikiandr/golang-todo-app"
	"github.com/pkg/errors"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) CreateItem(userId, listId int, curItem todo.Item) (int, error) {
	var userListsId int
	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id = $1 AND list_id=$2", usersListTable)

	err := r.db.Get(&userListsId, query, userId, listId)
	if err != nil {
		return 0, errors.New("no lists of this id for user")
	}

	query = fmt.Sprintf("INSERT INTO %s (list_id, title, description, done) VALUES ($1, $2, $3, $4) RETURNING id", itemsTable)

	var itemId int
	err = r.db.Get(&itemId, query, listId, curItem.Title, curItem.Description, curItem.Done)
	if err != nil {
		return 0, err
	}

	return itemId, nil
}

func (r *ItemPostgres) GetAllItems(userId, listId int) ([]todo.Item, error) {
	var items []todo.Item

	query := fmt.Sprintf("SELECT i.id, i.list_id, i.title, i.description, i.done "+
		"FROM (%s u INNER JOIN %s i ON u.list_id = i.list_id) "+
		"WHERE u.user_id = $1 AND i.list_id = $2", usersListTable, itemsTable)

	err := r.db.Select(&items, query, userId, listId)
	if err != nil {
		return items, err
	}

	return items, nil
}
