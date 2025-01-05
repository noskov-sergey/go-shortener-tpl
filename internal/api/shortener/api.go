package shortener

import (
	"github.com/go-chi/chi/v5"
)

//go:generate mockgen -source api.go -destination mocks/mocks.go -typed true service
type service interface {
	Create(string) (string, error)
	GetByID(string) (string, error)
}

type Implementation struct {
	service service
	chi.Router

	baseURL string
}

func New(service service, baseURL string) *Implementation {
	i := &Implementation{
		service: service,
		Router:  chi.NewRouter(),
		baseURL: baseURL,
	}
	i.Post("/", i.createHandler)
	i.Get("/{id}", i.getByIDHandler)

	return i
}
