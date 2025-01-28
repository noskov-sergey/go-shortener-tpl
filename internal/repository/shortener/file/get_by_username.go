package file

import (
	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

func (r *Repository) GetByUsername(username string) ([]model.Shortener, error) {
	var res []model.Shortener

	for k, v := range r.data {
		if len(v) != 2 {
			continue
		}
		if v[1] == username {
			res = append(res, model.Shortener{
				URL:      v[0],
				ShortURL: k,
			})
		}
	}

	return res, nil
}
