package shortener

import "github.ru/noskov-sergey/go-shortener-tpl/internal/model"

type Repo interface {
	Create(model.Shortener) error
	GetByID(string) (string, error)
	Ping() error
}

type service struct {
	repo Repo
}

func New(repo Repo) *service {
	return &service{
		repo: repo,
	}
}
