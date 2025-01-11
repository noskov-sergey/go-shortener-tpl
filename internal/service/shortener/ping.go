package shortener

func (s *service) Ping() error {
	if err := s.repo.Ping(); err != nil {
		return err
	}

	return nil
}
