package shortener

import (
	"github.com/go-chi/chi/v5"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/middleware"
)

type config struct {
	baseURL string
}

//go:generate mockgen -source api.go -destination mocks/mocks.go -typed true service
type service interface {
	Create(string) (string, error)
	GetByID(string) (string, error)
}

type Implementation struct {
	service service
	chi.Router

	cfg config
}

func New(service service, baseURL string) *Implementation {
	i := &Implementation{
		service: service,
		Router:  chi.NewRouter(),
		cfg: config{
			baseURL: baseURL,
		},
	}
	i.Use(middleware.WithLogging)
	i.Use(middleware.GzipMiddleware)
	i.Route("/", func(r chi.Router) {
		i.Post("/", i.createHandler)
		i.Get("/{id}", i.getByIDHandler)
		i.Route("/api", func(r chi.Router) {
			r.Route("/shorten", func(r chi.Router) {
				r.Post("/", i.shortenHandler)
			})
		})
	})
	i.Post("/", i.createHandler)
	i.Get("/{id}", i.getByIDHandler)

	return i
}
