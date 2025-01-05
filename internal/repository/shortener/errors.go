package shortener

import "errors"

var (
	ErrShortURLNotFound = errors.New("can't found by short url")
)
