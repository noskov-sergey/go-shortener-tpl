package shortener

import (
	"io"
	"log/slog"
	"net/http"
)

func (i *Implementation) createHandler(res http.ResponseWriter, req *http.Request) {
	log := i.log.With(slog.String("method", "createHandler"))
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Error("failed read body:", err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if string(body) == "" {
		log.Error("empty body:", err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	s, err := i.service.Create(string(body))
	if err != nil {
		log.Error("service create:", err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(i.cfg.baseURL + "/" + s))
}
