package shortener

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/converter"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

func (i *Implementation) createBatchHandler(res http.ResponseWriter, req *http.Request) {
	log := i.log.With(slog.String("method", "createHandler"))
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Error("failed read body:", slog.Any("err", err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	if len(body) == 0 {
		log.Error("empty body")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	var data []model.BatchRequest
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Error("unmarshal", slog.Any("err", err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	got, err := i.service.CreateBatch(converter.ToModelFromReq(data))
	if err != nil {
		log.Error("service create batch", slog.Any("err", err))
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("There are same urls ib db"))
		return
	}

	for x := range got {
		got[x].ShortURL = i.cfg.baseURL + "/" + got[x].ShortURL
	}

	r, err := json.MarshalIndent(converter.ToResFromModel(got), "", "    ")
	if err != nil {
		log.Error("marshal batch", slog.Any("err", err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(r)
}
