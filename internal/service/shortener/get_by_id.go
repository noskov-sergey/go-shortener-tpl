package shortener

import "fmt"

func (s *service) GetByID(shortUrl string) (string, error) {
	str, err := s.repo.GetByID(shortUrl)
	if err != nil {
		return "", fmt.Errorf("get by id: %w", err)
	}

	return str, nil
}
