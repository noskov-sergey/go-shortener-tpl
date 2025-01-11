package file

import (
	"fmt"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

func (r *repository) CreateBatchTx(data []model.Batch) error {
	for _, d := range data {
		r.data[d.ShortURL] = d.OriginalURL
		r.uuid++
	}

	if r.save {
		for _, d := range data {
			err := r.writeData(d.ShortURL, d.OriginalURL)
			if err != nil {
				return fmt.Errorf("error with write to file: %w", err)
			}
		}
	}

	return nil
}
