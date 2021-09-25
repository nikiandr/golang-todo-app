package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable     = "users"
	listsTable     = "todo_lists"
	usersListTable = "users_lists"
	itemsTable     = "todo_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresDBAuthString(auth string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(auth))

	if err != nil {
		return nil, err
	}

	return db, nil
}
