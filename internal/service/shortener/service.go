package shortener

type repo interface {
	Create(string) (string, error)
	GetByID(string) (string, error)
}

type service struct {
	repo repo
}

func New(repo repo) *service {
	return &service{
		repo: repo,
	}
}
