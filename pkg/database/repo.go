package database

import "database/sql"

type Repository struct {
	DB	*sql.DB
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{DB: database}
}