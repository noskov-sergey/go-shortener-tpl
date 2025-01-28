package file

import errors "github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener"

func (r *Repository) GetByID(shortURL string) (string, error) {
	if _, ok := r.data[shortURL]; !ok {
		return "", errors.ErrShortURLNotFound
	}
	return r.data[shortURL][0], nil
}
