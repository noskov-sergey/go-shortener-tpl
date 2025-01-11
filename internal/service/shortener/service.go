package shortener

type repo interface {
	Create(string) (string, error)
	GetByID(string) (string, error)
}

type dbRepo interface {
	Ping() error
}

type service struct {
	repo   repo
	dbRepo dbRepo
}

func New(repo repo, dbRepo dbRepo) *service {
	return &service{
		repo:   repo,
		dbRepo: dbRepo,
	}
}
