package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	shortenerApi "github.ru/noskov-sergey/go-shortener-tpl/internal/api/shortener"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/config"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener/file"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/service/shortener"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := config.New().Parse()

	f, err := os.OpenFile(cfg.File, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Warn("error open file", err)
		panic(err)
	}

	rep, err := file.New(f, cfg.Save == "file")
	if err != nil {
		log.Warn("error make repo", err)
		panic(err)
	}
	service := shortener.New(rep)
	imp := shortenerApi.New(service, cfg.BaseURL, log)

	log.Info(fmt.Sprintf("starting server on %s", cfg.URL))
	err = http.ListenAndServe(cfg.URL, imp)
	if err != nil {
		log.Warn("error starting http server", err)
		panic(err)
	}
}
