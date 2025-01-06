package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	shortenerApi "github.ru/noskov-sergey/go-shortener-tpl/internal/api/shortener"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/config"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener/memory"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/service/shortener"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := config.New().Parse()

	rep := memory.New()
	service := shortener.New(rep)
	imp := shortenerApi.New(service, cfg.BaseURL)

	log.Info(fmt.Sprintf("starting server on %s", cfg.URL))
	err := http.ListenAndServe(cfg.URL, imp)
	if err != nil {
		log.Warn("error starting http server", err)
		panic(err)
	}
}
