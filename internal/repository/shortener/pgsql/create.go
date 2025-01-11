package pgsql

import (
	"fmt"
	"time"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

func (r *Repository) Create(data model.Shortener) error {
	row := r.db.QueryRow(`
		INSERT INTO shortener (original_url,short_url,created_at)
		VALUES ($1, $2, $3) RETURNING id`,
		data.URL, data.ShortURL, time.Now())

	var id int
	err := row.Scan(&id)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
