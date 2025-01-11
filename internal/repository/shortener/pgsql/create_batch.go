package pgsql

import (
	"fmt"
	"time"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

func (r *Repository) CreateBatchTx(data []model.Batch) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("create transaction: %w", err)
	}

	for _, d := range data {
		_, err = tx.Exec(`
		INSERT INTO shortener (original_url,short_url,created_at)
		VALUES ($1, $2, $3) RETURNING id`,
			d.OriginalURL, d.ShortURL, time.Now())
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed exec tx: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
