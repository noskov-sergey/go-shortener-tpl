package shortener

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener"
)

func (i *Implementation) shortenHandler(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var data model.ShortenRequest
	err = json.Unmarshal(body, &data)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}

	s, err := i.service.Create(data.URL)
	if errors.Is(err, shortener.ErrNotUnique) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusConflict)

		got, err := json.Marshal(model.ToResponse(i.cfg.baseURL + "/" + s))
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		res.Write(got)
		return
	}

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	got, err := json.Marshal(model.ToResponse(i.cfg.baseURL + "/" + s))
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(got)
}
