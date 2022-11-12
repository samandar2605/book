package postgres

import (
	

	"github.com/jmoiron/sqlx"
)

type DBManager struct {
	db *sqlx.DB
}

func NewBookRepo(connStr *sqlx.DB) *DBManager {
	return &DBManager{db: connStr}
}
