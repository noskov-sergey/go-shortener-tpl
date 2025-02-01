package shortener

import "fmt"

func (s *service) GetByID(shortURL string) (string, error) {
	str, err := s.repo.GetByID(shortURL)
	if err != nil {
		return "", fmt.Errorf("get by id: %w", err)
	}

	return str.URL, nil
}
