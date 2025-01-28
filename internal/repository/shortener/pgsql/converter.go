package pgsql

import "github.ru/noskov-sergey/go-shortener-tpl/internal/model"

func toModel(s []Shortener) []model.Shortener {
	m := make([]model.Shortener, len(s))
	for i, v := range s {
		m[i] = model.Shortener{
			URL:      v.OriginalURL,
			ShortURL: v.ShortURL,
		}
	}

	return m
}
