package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

const (
	envServerAdress = "SERVER_ADDRESS"
	envBaseURL      = "BASE_URL"
)

type config struct {
	URL     string
	BaseURL string
}

func New() *config {
	return &config{}
}

func (c *config) Parse() *config {
	cfg := flag.String("c", ".env", "config file path")
	flag.StringVar(&c.URL, "a", ":8080", "address and port to run server")
	flag.StringVar(&c.BaseURL, "b", "http://localhost:8080", "base url to inform user")
	flag.Parse()

	godotenv.Load(*cfg)
	if URL := os.Getenv(envServerAdress); URL != "" {
		c.URL = URL
	}
	if BaseURL := os.Getenv(envBaseURL); BaseURL != "" {
		c.BaseURL = BaseURL
	}

	return c
}
