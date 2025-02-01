package file

import (
	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
	errors "github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener"
)

func (r *Repository) GetByID(shortURL string) (*model.Shortener, error) {
	if _, ok := r.data[shortURL]; !ok {
		return nil, errors.ErrShortURLNotFound
	}

	res := &model.Shortener{
		URL:      r.data[shortURL][0],
		ShortURL: shortURL,
	}
	return res, nil
}
