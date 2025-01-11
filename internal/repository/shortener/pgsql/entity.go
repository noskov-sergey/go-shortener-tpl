package pgsql

import "time"

type Shortener struct {
	ID          int64     `db:"id"`
	OriginalURL string    `db:"original_url"`
	ShortURL    string    `db:"short_url"`
	CreatedAt   time.Time `db:"created_at"`
}
