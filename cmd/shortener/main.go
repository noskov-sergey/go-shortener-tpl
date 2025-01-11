package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	shortenerApi "github.ru/noskov-sergey/go-shortener-tpl/internal/api/shortener"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/config"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener/file"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/service/shortener"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := config.New().Parse()

	f, err := os.OpenFile(cfg.File, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Error("error open file", slog.Any("err", err))
		panic(err)
	}

	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		log.Error("failed to connect to database:", slog.Any("err", err))
		panic(err)
	}
	defer db.Close()

	rep, err := file.New(f, cfg.Save == "file")
	if err != nil {
		log.Error("error make repo", slog.Any("err", err))
		panic(err)
	}
	service := shortener.New(rep)
	imp := shortenerApi.New(service, cfg.BaseURL, db, log)

	log.Info(fmt.Sprintf("starting server on %s", cfg.URL))
	err = http.ListenAndServe(cfg.URL, imp)
	if err != nil {
		log.Error("error starting http server", slog.Any("err", err))
		panic(err)
	}
}
