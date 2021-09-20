package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
