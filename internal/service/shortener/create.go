package shortener

import (
	"fmt"
	"math/rand"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

const (
	runeLen     = 52
	shortURLLen = 8
)

func (s *service) Create(url string) (string, error) {
	ru := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	str := make([]rune, shortURLLen)

	for i := range shortURLLen {
		str[i] = ru[rand.Intn(runeLen)]
	}

	data := model.Shortener{
		URL:      url,
		ShortURL: string(str),
	}

	err := s.repo.Create(data)

	if err != nil {
		return "", fmt.Errorf("create: %w", err)
	}

	return string(str), nil
}
