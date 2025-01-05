package main

import (
	"log"
	"net/http"

	shortenerApi "github.ru/noskov-sergey/go-shortener-tpl/internal/api/shortener"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/repository/shortener/memory"
	"github.ru/noskov-sergey/go-shortener-tpl/internal/service/shortener"
)

func main() {
	rep := memory.New()
	service := shortener.New(rep)
	imp := shortenerApi.New(service)

	err := http.ListenAndServe(":8080", imp)
	if err != nil {
		log.Fatal(err)
	}
}
