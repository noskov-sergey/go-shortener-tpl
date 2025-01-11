package converter

import "github.ru/noskov-sergey/go-shortener-tpl/internal/model"

func ToResFromModel(batch []model.Batch) []model.BatchResponse {
	var response []model.BatchResponse
	for _, b := range batch {
		response = append(response, model.BatchResponse{
			CorrelationID: b.CorrelationID,
			ShortURL:      b.ShortURL,
		})
	}

	return response
}

func ToModelFromReq(batch []model.BatchRequest) []model.Batch {
	var models []model.Batch
	for _, b := range batch {
		models = append(models, model.Batch{
			CorrelationID: b.CorrelationID,
			OriginalURL:   b.OriginalURL,
		})
	}
	return models
}
