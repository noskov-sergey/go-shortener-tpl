package pgsql

func (r *Repository) Ping() error {
	if err := r.db.Ping(); err != nil {
		return err
	}

	return nil
}
