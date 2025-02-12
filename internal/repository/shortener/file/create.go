package file

import (
	"encoding/json"
	"fmt"

	"github.ru/noskov-sergey/go-shortener-tpl/internal/model"
)

func (r *Repository) Create(shortener model.Shortener) error {
	r.data[shortener.ShortURL] = []string{shortener.URL, shortener.Username}
	r.uuid++

	if r.save {
		if err := r.writeData(shortener.ShortURL, shortener.URL, shortener.Username); err != nil {
			return fmt.Errorf("error with write to file: %w", err)
		}
	}
	return nil
}

func (r *Repository) writeData(str string, URL string, username string) error {
	line := Shorten{
		UUID:        r.uuid,
		ShortURL:    str,
		OriginalURL: URL,
		Username:    &username,
	}

	data, err := json.Marshal(&line)
	if err != nil {
		return fmt.Errorf("error marshalling shorten: %w", err)
	}

	_, err = r.writer.Write(data)
	if err != nil {
		return fmt.Errorf("error write shorten: %w", err)
	}

	err = r.writer.WriteByte('\n')
	if err != nil {
		return fmt.Errorf("error write /n: %w", err)
	}

	err = r.writer.Flush()
	if err != nil {
		return fmt.Errorf("error flush: %w", err)
	}
	return nil
}
