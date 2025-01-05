package shortener

import "net/http"

type service interface {
	Create(string) (string, error)
	GetByID(string) (string, error)
}

type Implementation struct {
	service service
	http.ServeMux
}

func New(service service) *Implementation {
	i := &Implementation{
		service:  service,
		ServeMux: *http.NewServeMux(),
	}
	i.HandleFunc("/", i.Create)
	i.HandleFunc("/{id}", i.GetByID)

	return i
}
