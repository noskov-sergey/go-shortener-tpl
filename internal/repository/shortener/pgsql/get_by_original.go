package pgsql

import "fmt"

func (r *Repository) GetByOriginal(URL string) (string, error) {
	query, err := r.db.Prepare(`SELECT id, original_url, short_url, created_at FROM shortener WHERE original_url = $1`)
	if err != nil {
		return "", fmt.Errorf("query prepare: %w", err)
	}

	var data Shortener

	err = query.QueryRow(URL).Scan(&data.ID, &data.OriginalURL, &data.ShortURL, &data.CreatedAt)
	if err != nil {
		return "", fmt.Errorf("row scan: %w", err)
	}

	return data.ShortURL, nil
}
