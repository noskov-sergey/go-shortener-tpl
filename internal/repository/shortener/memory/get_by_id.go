package memory

import (
	errors "github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener"
)

func (r *repository) GetByID(shortURL string) (string, error) {
	if _, ok := r.data[shortURL]; !ok {
		return "", errors.ErrShortUrlNotFound
	}
	return r.data[shortURL], nil
}
