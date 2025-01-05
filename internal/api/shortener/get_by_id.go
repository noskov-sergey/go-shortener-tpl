package shortener

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (i *Implementation) getByIDHandler(res http.ResponseWriter, req *http.Request) {
	shortURL := chi.URLParam(req, "id")
	if shortURL == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := i.service.GetByID(shortURL)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
