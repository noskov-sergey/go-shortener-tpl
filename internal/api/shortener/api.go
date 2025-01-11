package shortener

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/middleware"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

type config struct {
	baseURL string
}

//go:generate mockgen -source api.go -destination mocks/mocks.go -typed true service
type service interface {
	Create(string) (string, error)
	GetByID(string) (string, error)
	Ping() error
	CreateBatch([]model.Batch) ([]model.Batch, error)
}

type Implementation struct {
	service service
	chi.Router

	log *slog.Logger

	cfg config
}

func New(service service, baseURL string, log *slog.Logger) *Implementation {
	i := &Implementation{
		service: service,
		Router:  chi.NewRouter(),
		log:     log,
		cfg: config{
			baseURL: baseURL,
		},
	}
	i.Use(middleware.WithLogging)
	i.Use(middleware.GzipMiddleware)
	i.Route("/", func(r chi.Router) {
		i.Post("/", i.createHandler)
		i.Get("/{id}", i.getByIDHandler)
		i.Get("/ping", i.pingHandler)
		i.Route("/api", func(r chi.Router) {
			r.Route("/shorten", func(r chi.Router) {
				r.Post("/", i.shortenHandler)
				r.Post("/batch", i.createBatchHandler)
			})
		})
	})
	i.Post("/", i.createHandler)
	i.Get("/{id}", i.getByIDHandler)

	return i
}
