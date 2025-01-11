package pgsql

import (
	"fmt"
	"time"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

func (r *Repository) Create(data model.Shortener) error {
	got, err := r.db.Exec(`
		INSERT INTO shortener (original_url,short_url,created_at)
		VALUES ($1, $2, $3) RETURNING id`,
		data.URL, data.ShortURL, time.Now())

	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	row, _ := got.RowsAffected()
	if row == 0 {
		return fmt.Errorf("no row affected: %w", err)
	}

	return nil
}
