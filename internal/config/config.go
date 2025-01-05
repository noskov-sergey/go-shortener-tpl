package config

import (
	"flag"
)

type config struct {
	URL     string
	BaseURL string
}

func New() *config {
	return &config{}
}

func (c *config) ParseFlag() *config {
	flag.StringVar(&c.URL, "url", ":8080", "addres and port to run server")
	flag.StringVar(&c.BaseURL, "base", "http://localhost:8080/", "base url to inform user")
	flag.Parse()

	return c
}
