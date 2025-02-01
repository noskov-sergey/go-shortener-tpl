package shortener

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func (i *Implementation) deleteHandler(res http.ResponseWriter, req *http.Request) {
	log := i.log.With(slog.String("method", "deleteHandler"))
	user := req.Header.Get(AuthLogin)

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

	var data []string
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Error("unmarshal", slog.Any("err", err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	go func() {
		err = i.service.Delete(ctx, user, data)
		if err != nil {
			log.Error("service delete", slog.Any("err", err))
		}
	}()

	res.WriteHeader(http.StatusAccepted)
}
