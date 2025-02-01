package pgsql

import (
	"fmt"
)

func (r *Repository) MarkDelete(data []string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("begin: %w", err)
	}

	for _, shortUrl := range data {
		_, err = tx.Exec(`
		UPDATE shortener SET is_deleted = TRUE
		WHERE short_url = $1`,
			shortUrl)
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
