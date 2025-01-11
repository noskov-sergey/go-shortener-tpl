package shortener

import (
	"log/slog"
	"net/http"
)

func (i *Implementation) pingHandler(res http.ResponseWriter, req *http.Request) {
	log := i.log.With(slog.String("method", "pingHandler"))

	if err := i.DB.Ping(); err != nil {
		log.Error("failed to ping database", slog.Any("error", err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}
