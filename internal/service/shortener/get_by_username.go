package shortener

import (
	"fmt"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

func (s *service) GetByUsername(username string) ([]model.Shortener, error) {
	got, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("get by username: %w", err)
	}

	return got, nil
}
