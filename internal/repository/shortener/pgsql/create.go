package pgsql

import (
	"fmt"
	"time"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener"
)

func (r *Repository) Create(data model.Shortener) error {
	got, err := r.db.Exec(`
		INSERT INTO shortener (original_url,short_url,created_at, username)
		VALUES ($1, $2, $3, $4) ON CONFLICT (original_url) DO NOTHING`,
		data.URL, data.ShortURL, time.Now(), data.Username)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	row, _ := got.RowsAffected()
	if row == 0 {
		return shortener.ErrNotUnique
	}

	return nil
}
