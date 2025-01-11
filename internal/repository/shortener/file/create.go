package file

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

const (
	runeLen     = 52
	shortURLLen = 8
)

func (r *repository) Create(URL string) (string, error) {
	ru := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	str := make([]rune, shortURLLen)

	for i := range shortURLLen {
		str[i] = ru[rand.Intn(runeLen)]
	}

	r.data[string(str)] = URL
	r.uuid++

	if r.save {
		if err := r.writeData(str, URL); err != nil {
			return "", fmt.Errorf("error with write to file: %w", err)
		}
	}
	return string(str), nil
}

func (r *repository) writeData(str []rune, URL string) error {
	line := Shorten{
		UUID:        r.uuid,
		ShortURL:    string(str),
		OriginalURL: URL,
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
