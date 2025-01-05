package shortener

import "fmt"

func (s *service) Create(url string) (string, error) {
	str, err := s.repo.Create(url)
	if err != nil {
		return "", fmt.Errorf("create: %w", err)
	}

	return str, nil
}
