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
		return nil, err
	}

	return items, nil
}

func (r *ItemPostgres) GetItemById(userId, listId, itemId int) (todo.Item, error) {
	var item todo.Item

	query := fmt.Sprintf("SELECT i.id, i.list_id, i.title, i.description, i.done "+
		"FROM (%s u INNER JOIN %s i ON u.list_id = i.list_id) "+
		"WHERE u.user_id = $1 AND i.list_id = $2 AND i.id = $3", usersListTable, itemsTable)
	err := r.db.Get(&item, query, userId, listId, itemId)
	if err != nil {
		return todo.Item{}, err
	}
	return item, nil
}

func (r *ItemPostgres) UpdateItem(update todo.Item, userId, listId, itemId int) error {
	var resUserId int

	query := fmt.Sprintf("UPDATE %s SET title = $1, description = $2, done = $3 "+
		"WHERE id IN (SELECT i.id FROM (%s u INNER JOIN %s i ON u.list_id = i.list_id) "+
		"WHERE u.user_id = $4 AND i.list_id = $5 AND i.id = $6) RETURNING id", itemsTable, usersListTable, itemsTable)
	err := r.db.Get(&resUserId, query, update.Title, update.Description, update.Done, userId, listId, itemId)
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemPostgres) DeleteItem(userId, listId, itemId int) error {
	var resUserId int

	query := fmt.Sprintf("DELETE FROM %s WHERE id IN "+
		"(SELECT i.id FROM (%s u INNER JOIN %s i ON u.list_id = i.list_id) "+
		"WHERE u.user_id = $1 AND i.list_id = $2 AND i.id = $3) RETURNING id", itemsTable, usersListTable, itemsTable)
	err := r.db.Get(&resUserId, query, userId, listId, itemId)
	if err != nil {
		return err
	}

	return nil
}
