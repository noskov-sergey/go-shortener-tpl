package shortener

import "github.ru/noskov-sergey/go-shortener-tpl/internal/model"

type Repo interface {
	Create(model.Shortener) error
	GetByID(string) (*model.Shortener, error)
	Ping() error
	CreateBatchTx([]model.Batch) error
	GetByOriginal(string) (string, error)
	GetByUsername(string) ([]model.Shortener, error)
	MarkDelete(data []string) error
}

type service struct {
	repo Repo
}

func New(repo Repo) *service {
	return &service{
		repo: repo,
	}
}
