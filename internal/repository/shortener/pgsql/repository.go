package pgsql

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewShortenRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
