package shortener

import (
	"errors"
	"fmt"
	"math/rand"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener"
)

const (
	runeLen     = 52
	shortURLLen = 8
)

func (s *service) Create(data model.Shortener) (string, error) {
	data.ShortURL = generateShortURL()

	err := s.repo.Create(data)
	if err != nil {
		if errors.Is(err, shortener.ErrNotUnique) {
			got, errGet := s.repo.GetByOriginal(data.URL)
			if errGet != nil {
				return "", errGet
			}
			return got, shortener.ErrNotUnique
		}
		return "", fmt.Errorf("create: %w", err)
	}

	return data.ShortURL, nil
}

func generateShortURL() string {
	ru := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	str := make([]rune, shortURLLen)

	for i := range shortURLLen {
		str[i] = ru[rand.Intn(runeLen)]
	}

	return string(str)
}
