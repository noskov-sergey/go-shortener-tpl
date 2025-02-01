package pgsql

import (
	"fmt"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener"
)

func (r *Repository) GetByID(shortURL string) (*model.Shortener, error) {
	query, err := r.db.Prepare(`SELECT id, original_url, short_url, created_at, username, is_deleted FROM shortener WHERE short_url = $1`)
	if err != nil {
		return nil, fmt.Errorf("query prepare: %w", err)
	}

	var data Shortener

	err = query.QueryRow(shortURL).Scan(&data.ID, &data.OriginalURL, &data.ShortURL, &data.CreatedAt, &data.Username, &data.Deleted)
	if err != nil {
		return nil, fmt.Errorf("row scan: %w", err)
	}

	if data.Deleted {
		return nil, shortener.ErrDeleted
	}

	res := &model.Shortener{
		URL:      data.OriginalURL,
		ShortURL: data.ShortURL,
		Username: *data.Username,
	}

	return res, nil
}
