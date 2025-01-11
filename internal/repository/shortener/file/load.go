package file

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (r *repository) load() error {
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

		r.uuid++
		r.data[data.ShortURL] = data.OriginalURL
	}

	return nil
}
