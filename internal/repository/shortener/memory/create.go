package memory

import "github.ru/noskov-sergey/go-shortener-tpl/internal/model"

func (r *repository) Create(data model.Shortener) error {
	r.data[data.ShortURL] = data.URL

	return nil
}
