package shortener

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener"
)

func (i *Implementation) getByIDHandler(res http.ResponseWriter, req *http.Request) {
	shortURL := chi.URLParam(req, "id")
	if shortURL == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := i.service.GetByID(shortURL)
	if err != nil {
		if errors.Is(err, shortener.ErrDeleted) {
			res.WriteHeader(http.StatusGone)
			return
		}
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
