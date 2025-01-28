package file

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (r *Repository) load() error {
	for r.scanner.Scan() {
		line := r.scanner.Bytes()
		if len(line) == 0 {
			return errors.New("file is clear")
		}

		data := Shorten{}
		err := json.Unmarshal(line, &data)
		if err != nil {
			return fmt.Errorf("json unmarshal error: %v", err)
		}

		var user string
		if data.Username == nil {
			user = ""
		} else {
			user = *data.Username
		}

		r.uuid++
		r.data[data.ShortURL] = []string{data.OriginalURL, user}
	}

	return nil
}
