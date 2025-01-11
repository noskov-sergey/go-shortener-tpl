package shortener

import (
	"fmt"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

func (s *service) CreateBatch(data []model.Batch) ([]model.Batch, error) {
	var goodData []model.Batch
	for _, d := range data {
		if d.OriginalURL == "" || d.CorrelationID == "" {
			continue
		}

		good := model.Batch{
			CorrelationID: d.CorrelationID,
			ShortURL:      generateShortURL(),
			OriginalURL:   d.OriginalURL,
		}

		goodData = append(goodData, good)
	}

	err := s.repo.CreateBatchTx(goodData)
	if err != nil {
		return nil, fmt.Errorf("create batch tx: %w", err)
	}

	return goodData, nil
}
