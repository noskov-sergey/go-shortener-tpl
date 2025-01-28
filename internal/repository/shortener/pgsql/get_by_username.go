package pgsql

import (
	"fmt"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

func (r *Repository) GetByUsername(username string) ([]model.Shortener, error) {
	query, err := r.db.Prepare(`SELECT id, original_url, short_url, created_at, username FROM shortener WHERE username = $1`)
	if err != nil {
		return nil, fmt.Errorf("query prepare: %w", err)
	}

	var data []Shortener

	rows, err := query.Query(username)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	for rows.Next() {
		var d Shortener
		err = rows.Scan(&d.ID, &d.OriginalURL, &d.ShortURL, &d.CreatedAt, &d.Username)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		data = append(data, d)
	}

	return toModel(data), nil
}
