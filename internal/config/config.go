package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

const (
	envServerAdress = "SERVER_ADDRESS"
	envBaseURL      = "BASE_URL"
	envFile         = "FILE_STORAGE_PATH"
	envDSN          = "DATABASE_DSN"

	envFileDefault = "./tmp/short-url-db.json"

	repoFileValue = "file"
	repoDBValue   = "db"
)

type config struct {
	URL     string
	BaseURL string
	File    string
	Save    string
	DSN     string
}

func New() *config {
	return &config{}
}

func (c *config) Parse() *config {
	c.File = envFileDefault
	cfg := flag.String("c", ".env", "config file path")
	flag.StringVar(&c.URL, "a", ":8080", "address and port to run server")
	flag.StringVar(&c.BaseURL, "b", "http://localhost:8080", "base url to inform user")
	//flag.StringVar(&c.File, "f", envFileDefault, "base filepath to backup data")
	flag.Func("f", "base filepath to backup data", func(flagValue string) error {
		if flagValue != "" {
			c.File = flagValue
			c.Save = repoFileValue
		}
		return nil
	})
	flag.Func("d", "database connection string", func(flagValue string) error {
		if flagValue != "" {
			c.DSN = flagValue
			c.Save = repoDBValue
		}
		return nil
	})

	flag.Parse()

	godotenv.Load(*cfg)
	if URL := os.Getenv(envServerAdress); URL != "" {
		c.URL = URL
	}

	if BaseURL := os.Getenv(envBaseURL); BaseURL != "" {
		c.BaseURL = BaseURL
	}

	if File := os.Getenv(envFile); File != "" {
		c.File = File
		c.Save = repoFileValue
	}

	if DSN := os.Getenv(envDSN); DSN != "" {
		c.DSN = DSN
		c.Save = repoDBValue
	}

	return c
}
