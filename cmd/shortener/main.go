package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	shortenerApi "github.ru/noskov-sergey/go-shortener-tpl/internal/api/shortener"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener/memory"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/service/shortener"
)

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Error initializing logger")
	}

	rep := memory.New()
	service := shortener.New(rep)
	imp := shortenerApi.New(service)

	log.Info("starting server on :8080")
	err = http.ListenAndServe(":8080", imp)
	if err != nil {
		log.Fatal("error starting http server", zap.Error(err))
	}
}
