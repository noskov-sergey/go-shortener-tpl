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
}

func New(service service) *Implementation {
	i := &Implementation{
		service: service,
		Router:  chi.NewRouter(),
	}
	i.Post("/", i.createHandler)
	i.Get("/{id}", i.getByIDHandler)

	return i
}
