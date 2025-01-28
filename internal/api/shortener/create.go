package shortener

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener"
)

func (i *Implementation) createHandler(res http.ResponseWriter, req *http.Request) {
	log := i.log.With(slog.String("method", "createHandler"))
	user := req.Header.Get(AuthLogin)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Error("failed read body:", slog.Any("err", err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if string(body) == "" {
		log.Error("empty body")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	data := model.Shortener{
		URL:      string(body),
		Username: user,
	}

	s, err := i.service.Create(data)
	if errors.Is(err, shortener.ErrNotUnique) {
		res.WriteHeader(http.StatusConflict)
		res.Write([]byte(i.cfg.baseURL + "/" + s))
		return
	}

	if err != nil {
		log.Error("service create:", slog.Any("err", err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(i.cfg.baseURL + "/" + s))
}
