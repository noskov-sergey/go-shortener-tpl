package shortener

func (s *service) Ping() error {
	if err := s.dbRepo.Ping(); err != nil {
		return err
	}

	return nil
}
