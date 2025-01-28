package shortener

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

func (i *Implementation) getByUsernameHandler(res http.ResponseWriter, req *http.Request) {
	log := i.log.With(slog.String("method", "getByUsername"))
	user := req.Header.Get(AuthLogin)

	data, err := i.service.GetByUsername(user)
	if err != nil {
		log.Error("failed read body:", slog.Any("error", err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(data) == 0 {
		log.Error("no data:", slog.Any("error", err))
		res.WriteHeader(http.StatusNoContent)
		return
	}

	for x := range data {
		data[x].ShortURL = i.cfg.baseURL + "/" + data[x].ShortURL
	}

	r, err := json.MarshalIndent(model.ToUserResponse(data), "", "    ")
	if err != nil {
		log.Error("marshal", slog.Any("err", err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}
